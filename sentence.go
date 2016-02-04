package jk

import (
	"errors"
	"strings"
)

//Morphemes is a slice of Morpheme
type Morphemes []*Morpheme

//Sentence includes elements of a sentence
type Sentence struct {
	Morphemes
	Bunsetsus    DependencyInfos
	BasicPhrases DependencyInfos
	comment      string
}

//NewSentence creats a sentence with the given text
func NewSentence(lines []string) (*Sentence, error) {
	sent := new(Sentence)
	sent.Bunsetsus = DependencyInfos{}
	sent.BasicPhrases = DependencyInfos{}

	for _, line := range lines {

		if strings.HasPrefix(line, "#") {
			sent.comment += line
		} else if strings.HasPrefix(line, "EOS") {
			break
		} else if strings.HasPrefix(line, "@") {
			if len(line) < 2 {
				return sent, errors.New("The length less than 2")
			}
			m, err := NewMorpheme(line[2:])
			if err != nil {
				return sent, err
			}

			if len(sent.Morphemes) == 0 {
				return sent, errors.New("@ comes before some morpheme")
			}
			doukeis := &(sent.Morphemes[len(sent.Morphemes)-1].Doukeis)
			*doukeis = append(*doukeis, m)
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
		} else {
			m, err := NewMorpheme(line)
			if err != nil {
				return sent, err
			}
			sent.Morphemes = append(sent.Morphemes, m)
		}
	}

	return sent, nil
}

//GetMorphemes returns morpheme of the sentence
func (sent *Sentence) GetMorphemes() Morphemes {
	return sent.Morphemes
}

//Len returns the number of the morphemes
func (sent *Sentence) Len() int {
	return len(sent.Morphemes)
}
