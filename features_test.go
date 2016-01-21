package jk

import (
	"testing"
)

func TestGetFeatures(t *testing.T) {
	sysouts := []Features{
		getFeatures(sampleFeature, '>', 1),
		getFeatures(sampleFeature2, '|', 0),
	}
	gold := Features{
		`代表表記`:    `構文/こうぶん`,
		`カテゴリ`:    `抽象物`,
		`正規化代表表記`: `構文/こうぶん`,
		`漢字`:      ``,
	}

	for i, sysout := range sysouts {
		t.Logf("%v", sysout)
		if len(sysout) != len(gold) {
			t.Errorf("Size error: Expected %d but %d (in loop %d)", len(gold), len(sysout), i)
		}
		for k, gv := range gold {
			sysv, ok := sysout[k]
			if !ok {
				t.Errorf("For key [%v], expected [%v] but got nothing", k, gv)
			} else if gv != sysv {
				t.Errorf("For key [%v], expected [%v] but got [%v]", k, gv, sysv)
			}
		}
	}

}
