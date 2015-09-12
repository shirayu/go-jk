package jk

type JumanSocketClient struct {
	*SocketClient
}

func NewJumanSocketClient(address string) (*JumanSocketClient, error) {
	client, err := NewSocketClient(address, "RUN -e2\n")
	if err != nil {
		return nil, err
	}
	self := &JumanSocketClient{client}
	return self, err
}

func (self *JumanSocketClient) Parse(query string) (*Sentence, error) {
	lines, err := self.RawParse(query)
	if err != nil {
		return nil, err
	}
	s, err := NewSentence(lines)
	return s, err
}
