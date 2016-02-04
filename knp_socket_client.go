package jk

import "strings"

//KnpSocketClient is a client to a communicete knp server
type KnpSocketClient struct {
	*SocketClient
}

//NewKnpSocketClient creates a new KnpSocketClient
func NewKnpSocketClient(address string) (*KnpSocketClient, error) {
	client, err := NewSocketClient(address, "RUN -normal -tab\n")
	if err != nil {
		return nil, err
	}
	knpsc := &KnpSocketClient{client}
	return knpsc, err
}

//Parse returns a Sentence to parse the given sentence
func (knpsc *KnpSocketClient) Parse(jumanLines string) (*Sentence, error) {
	lines, err := knpsc.RawParse(jumanLines)
	if err != nil {
		return nil, err
	}
	s, err := NewSentence(strings.Split(lines, "\n"))
	return s, err
}
