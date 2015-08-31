package jk

import (
	"errors"
	"strings"
)

type Morphemes []*Morpheme

type Sentence struct {
	Morphemes
	Bunsetsus    DependencyInfos
	BasicPhrases DependencyInfos
	comment      string
}

func NewSentence(lines string) (*Sentence, error) {
	self := new(Sentence)
	self.Bunsetsus = DependencyInfos{}
	self.BasicPhrases = DependencyInfos{}

	for _, line := range strings.Split(lines, "\n") {

		if strings.HasPrefix(line, "#") {
			self.comment += line
		} else if strings.HasPrefix(line, "EOS") {
			break
		} else if strings.HasPrefix(line, "@") {
			if len(line) < 2 {
				return self, errors.New("The length less than 2")
			}
			m, err := NewMorpheme(line[2:])
			if err != nil {
				return self, err
			}

			if len(self.Morphemes) == 0 {
				return self, errors.New("@ comes before some morpheme")
			} else {
				doukeis := &(self.Morphemes[len(self.Morphemes)-1].Doukeis)
				*doukeis = append(*doukeis, m)
			}
		} else if strings.HasPrefix(line, "* ") {
			di, err := NewDependencyInfo(line)
			if err != nil {
				return self, err
			}
			self.Bunsetsus = append(self.Bunsetsus, di)
		} else if strings.HasPrefix(line, "+ ") {
			di, err := NewDependencyInfo(line)
			if err != nil {
				return self, err
			}
			self.BasicPhrases = append(self.BasicPhrases, di)
		} else {
			m, err := NewMorpheme(line)
			if err != nil {
				return self, err
			}
			self.Morphemes = append(self.Morphemes, m)
		}
	}

	return self, nil
}

func (self *Sentence) GetMorphemes() Morphemes {
	return self.Morphemes
}

func (self *Sentence) Len() int {
	return len(self.Morphemes)
}
