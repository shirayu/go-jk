package jk

import (
	"strings"
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
	if c := strings.Count(ret_lines, "\n"); c != 3 {
		t.Errorf("expceted length is 3 but %d", c)
	}

	s, err := juman.Parse("パンが食べられる")
	if err != nil {
		t.Error("Error to parse [%v]", err)
	}
	if s.Len() != 4 {
		t.Errorf("expceted length is 4 but %d", s.Len())
	}

}
