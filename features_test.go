package jk

import (
	"testing"
)

func TestGetFeatures(t *testing.T) {
	kvs := GetFeatures(`<代表表記:構文/こうぶん><カテゴリ:抽象物><正規化代表表記:構文/こうぶん><漢字>`)
	gold := Features{
		`代表表記`:    `構文/こうぶん`,
		`カテゴリ`:    `抽象物`,
		`正規化代表表記`: `構文/こうぶん`,
		`漢字`:      ``,
	}

	if len(kvs) != len(gold) {
		t.Errorf("Size error")
	}
	for k, gv := range gold {
		sysv, ok := kvs[k]
		if !ok {
			t.Errorf("For key [%v], expected [%v] but got nothing", k, gv)
		} else if gv != sysv {
			t.Errorf("For key [%v], expected [%v] but got [%v]", k, gv, sysv)
		}
	}

}
