package jk

import (
	"testing"
)

func TestJuman(t *testing.T) {
	juman, err := NewJumanClient("localhost:32000")
	if err != nil {
		t.Fatal("Error to open the juman socket: ", err)
	}

	ret_lines, err := juman.RawParse("パンが食べられる")
	if err != nil {
		t.Error("Error to parse [%v]", err)
	}
	if len(ret_lines) != 4 {
		t.Errorf("expceted length is 4 but %d", len(ret_lines))
	}

	s, err := juman.Parse("パンが食べられる")
	if err != nil {
		t.Error("Error to parse [%v]", err)
	}
	if s.Len() != 4 {
		t.Errorf("expceted length is 4 but %d", s.Len())
	}

}
