package jk

//Juman is a client to execute juman command
type Juman struct {
	*CommandClient
}

//NewJuman creates a new Juman
func NewJuman(path string) (*Juman, error) {
	client, err := NewCommandClient(path)
	if err != nil {
		return nil, err
	}
	jmn := &Juman{client}
	return jmn, err
}

//Parse returns a Sentence to parse the given sentence
func (jmn *Juman) Parse(query string) (*Sentence, error) {
	lines, err := jmn.RawParse(query)
	if err != nil {
		return nil, err
	}
	s, err := NewSentence(lines)
	return s, err
}
