package jk

type Juman struct {
	*CommandClient
}

func NewJuman(path string) (*Juman, error) {
	client, err := NewCommandClient(path)
	if err != nil {
		return nil, err
	}
	self := &Juman{client}
	return self, err
}

func (self *Juman) Parse(query string) (*Sentence, error) {
	lines, err := self.RawParse(query)
	if err != nil {
		return nil, err
	}
	s, err := NewSentence(lines)
	return s, err
}
