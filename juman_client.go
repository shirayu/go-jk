package jk

type JumanClient struct {
	*Client
}

func NewJumanClient(address string) (*JumanClient, error) {
	client, err := NewClient(address, "RUN -e2\n")
	if err != nil {
		return nil, err
	}
	self := &JumanClient{client}
	return self, err
}

func (self *JumanClient) Parse(query string) (*Sentence, error) {
	lines, err := self.RawParse(query)
	if err != nil {
		return nil, err
	}
	s, err := NewSentence(lines)
	return s, err
}
