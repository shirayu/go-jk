package jk

//Knp is a client to execute knp command
type Knp struct {
	*CommandClient
}

//NewKnp creates a new Knp
func NewKnp(path string, options ...string) (*Knp, error) {
	t := make([]string, len(options)+1)
	for i, opt := range options {
		t[i] = opt
	}
	t[len(options)] = "-tab"

	client, err := NewCommandClient(path, t...)
	if err != nil {
		return nil, err
	}
	knp := &Knp{client}
	return knp, err
}

//Parse returns a Sentence to parse the given sentence
func (knp *Knp) Parse(query string) (*Sentence, error) {
	lines, err := knp.RawParse(query)
	if err != nil {
		return nil, err
	}
	s, err := NewSentence(lines)
	return s, err
}
