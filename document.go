package jk

import "bufio"

//Sentences a slice of Sentence
type Sentences []*Sentence

//Document is a set of sentences
type Document struct {
	Sentences Sentences
}

//GetSentenceIndex returns the index by SID
func (doc *Document) GetSentenceIndex(sid string) int {
	for i, s := range doc.Sentences {
		if s.ID == sid {
			return i
		}
	}
	return -1
}

//NewDocument creates Document
func NewDocument(scanner *bufio.Scanner) (*Document, error) {
	lines := []string{}
	doc := Document{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
		if line == "EOS" {
			sent, err := NewSentence(lines)
			if err != nil {
				return &doc, err
			}
			doc.Sentences = append(doc.Sentences, sent)
			lines = []string{}
		}
	}
	return &doc, scanner.Err()
}
