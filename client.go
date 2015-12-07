package jk

//Client parses the given plain texts
type Client interface {
	RawParse(query string) (string, error)
	Parse(query string) (*Sentence, error)
}
