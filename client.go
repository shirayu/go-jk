package jk

type Client interface {
	RawParse(query string) (string, error)
	Parse(query string) (*Sentence, error)
}
