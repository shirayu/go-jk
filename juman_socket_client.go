package jk

//JumanSocketClient is a client to a communicete juman server
type JumanSocketClient struct {
	*SocketClient
}

//NewJumanSocketClient creates a new JumanSocketClient
func NewJumanSocketClient(address string) (*JumanSocketClient, error) {
	client, err := NewSocketClient(address, "RUN -e2\n")
	if err != nil {
		return nil, err
	}
	jsc := &JumanSocketClient{client}
	return jsc, err
}

//Parse returns a Sentence to parse the given sentence
func (jsc *JumanSocketClient) Parse(query string) (*Sentence, error) {
	lines, err := jsc.RawParse(query)
	if err != nil {
		return nil, err
	}
	s, err := NewSentence(lines)
	return s, err
}
