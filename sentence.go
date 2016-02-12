package jk

import (
	"errors"
	"strings"
	"unicode/utf8"
)

//Morphemes is a slice of Morpheme
type Morphemes []*Morpheme

//Sentence includes elements of a sentence
type Sentence struct {
	Morphemes
	ID                        string
	Bunsetsus                 DependencyInfos
	BasicPhrases              DependencyInfos
	comment                   string
	MorphemePositions         []int
	BasicPhrasePositions      []int
	BasicPhraseMorphemeIndexs []int
}

func (sent *Sentence) setComment(line string) {
	sent.comment += line
	if strings.HasPrefix(line, "# S-ID:") {
		tail := line[7:]
		end := strings.Index(tail, " ")
		if end < 0 {
			sent.ID = tail
		} else {
			sent.ID = tail[:end]
		}
	}
}

func (sent *Sentence) setDoukei(line string) error {
	if len(line) < 2 {
		return errors.New("The length less than 2")
	}
	m, err := NewMorpheme(line[2:])
	if err != nil {
		return err
	}

	if len(sent.Morphemes) == 0 {
		return errors.New("@ comes before some morpheme")
	}
	doukeis := &(sent.Morphemes[len(sent.Morphemes)-1].Doukeis)
	*doukeis = append(*doukeis, m)
	return nil
}

//NewSentence creats a sentence with the given text
func NewSentence(lines []string) (*Sentence, error) {
	sent := new(Sentence)
	sent.Bunsetsus = DependencyInfos{}
	sent.BasicPhrases = DependencyInfos{}
	sent.MorphemePositions = []int{0}
	sent.BasicPhrasePositions = []int{}

	length := 0
	for _, line := range lines {

		if strings.HasPrefix(line, "#") {
			sent.setComment(line)
		} else if strings.HasPrefix(line, "EOS") {
			break
		} else if strings.HasPrefix(line, "@") {
			if err := sent.setDoukei(line); err != nil {
				return sent, errors.New("The length less than 2")
			}
		} else if strings.HasPrefix(line, "* ") {
			di, err := NewDependencyInfo(line)
			if err != nil {
				return sent, err
			}
			sent.Bunsetsus = append(sent.Bunsetsus, di)
		} else if strings.HasPrefix(line, "+ ") {
			di, err := NewDependencyInfo(line)
			if err != nil {
				return sent, err
			}
			sent.BasicPhrases = append(sent.BasicPhrases, di)
			sent.BasicPhrasePositions = append(sent.BasicPhrasePositions, length)
			sent.BasicPhraseMorphemeIndexs = append(sent.BasicPhraseMorphemeIndexs, len(sent.Morphemes))
		} else {
			m, err := NewMorpheme(line)
			if err != nil {
				return sent, err
			}
			sent.Morphemes = append(sent.Morphemes, m)
			length += utf8.RuneCountInString(m.Surface)
			sent.MorphemePositions = append(sent.MorphemePositions, length)
		}
	}
	sent.BasicPhrasePositions = append(sent.BasicPhrasePositions, length)
	sent.BasicPhraseMorphemeIndexs = append(sent.BasicPhraseMorphemeIndexs, len(sent.Morphemes))

	return sent, nil
}

//GetMorphemes returns morpheme of the sentence
func (sent *Sentence) GetMorphemes(bpIndex int) Morphemes {
	if bpIndex < 0 || bpIndex >= len(sent.BasicPhrasePositions) {
		return nil
	}
	start := sent.BasicPhraseMorphemeIndexs[bpIndex]
	end := sent.BasicPhraseMorphemeIndexs[bpIndex+1]
	return sent.Morphemes[start:end]
}

//Len returns the number of the morphemes
func (sent *Sentence) Len() int {
	return len(sent.Morphemes)
}
