package jk

type KnpSocketClient struct {
	*SocketClient
}

func NewKnpSocketClient(address string) (*KnpSocketClient, error) {
	client, err := NewSocketClient(address, "RUN -normal -tab\n")
	if err != nil {
		return nil, err
	}
	self := &KnpSocketClient{client}
	return self, err
}

func (self *KnpSocketClient) Parse(juman_lines string) (*Sentence, error) {
	lines, err := self.RawParse(juman_lines)
	if err != nil {
		return nil, err
	}
	s, err := NewSentence(lines)
	return s, err
}
