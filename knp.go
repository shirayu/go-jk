package jk

//Knp is a client to execute knp command
type Knp struct {
	*CommandClient
}

//NewKnp creates a new Knp
func NewKnp(path string, options ...string) (*Knp, error) {
	client, err := NewCommandClient(path, "-tab")
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
