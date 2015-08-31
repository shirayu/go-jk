package jk

type KnpClient struct {
	*Client
}

func NewKnpClient(address string) (*KnpClient, error) {
	client, err := NewClient(address, "RUN -tab\n")
	if err != nil {
		return nil, err
	}
	self := &KnpClient{client}
	return self, err
}

func (self *KnpClient) Parse(juman_lines string) (*Sentence, error) {
	lines, err := self.RawParse(juman_lines)
	if err != nil {
		return nil, err
	}
	s, err := NewSentence(lines)
	return s, err
}
