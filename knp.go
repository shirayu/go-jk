package jk

type Knp struct {
	*CommandClient
}

func NewKnp(path string, options ...string) (*Knp, error) {
	client, err := NewCommandClient(path, "-tab")
	if err != nil {
		return nil, err
	}
	self := &Knp{client}
	return self, err
}

func (self *Knp) Parse(query string) (*Sentence, error) {
	lines, err := self.RawParse(query)
	if err != nil {
		return nil, err
	}
	s, err := NewSentence(lines)
	return s, err
}
