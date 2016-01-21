package jk

import (
	"testing"
)

func TestParsePas(t *testing.T) {
	pas, err := NewPas(samplePasResult, true)
	if err != nil {
		t.Errorf("Expected not nil")
	}

	goldCfid := "分/ふん:判1"
	if pas.Cfid != goldCfid {
		t.Errorf("Expected [%v] but got [%v]", goldCfid, pas.Cfid)
	}

	f2args := pas.Args
	if len(f2args) != 2 {
		t.Errorf("The length expected [%v] but got [%v]", 2, len(f2args))
	} else if len(f2args["デ"]) != 2 {
		t.Errorf("The length expected [%v] but got [%v]", 2, len(f2args["デ"]))
	} else if f2args["デ"][0].Sid != "14" {
		t.Errorf("Error in parse")
	} else if f2args["デ"][0].Tid != 1 {
		t.Errorf("Error in parse")
	} else if f2args["デ"][0].Rep != "車" {
		t.Errorf("Error in parse")
	} else if f2args["デ"][1].Sid != "17" {
		t.Errorf("Error in parse")
	} else if f2args["デ"][1].Tid != 7 {
		t.Errorf("Error in parse")
	} else if f2args["デ"][1].Rep != "徒歩" {
		t.Errorf("Error in parse")
	} else if len(f2args["ヨリ"]) != 1 {
		t.Errorf("The length expected [%v] but got [%v]", 1, len(f2args["ヨリ"]))
	} else if len(f2args["ガ"]) != 0 {
		t.Errorf("Error in parse")
	}
	//         self.assertEqual(f2.rels, None)

}
