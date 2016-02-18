package jk

import (
	"reflect"
	"testing"
)

func TestMorpheme(t *testing.T) {
	line := "探して さがして 探す 動詞 2 * 0 子音動詞サ行 5 タ系連用テ形 14 \"代表表記:探す/さがす\""
	m, err := NewMorpheme(line)

	if err != nil {
		t.Fatal(err)
	}

	if m.Surface != "探して" {
		t.Errorf("Midashi Error\n")
	} else if m.CFormID != 14 {
		t.Errorf("Katsuyou2ID Error\n")
	} else if m.Seminfo != "代表表記:探す/さがす" {
		t.Errorf("Seminfo_id Error\n")
	} else if m.Rep != "探す/さがす" {
		t.Errorf("Rep Error\n")
	}
}

func TestMorpheme2(t *testing.T) {
	line := `を を を 助詞 9 格助詞 1 * 0 * 0 NIL <かな漢字><ひらがな><付属>`
	m, err := NewMorpheme(line)

	if err != nil {
		t.Fatal(err)
	}

	if m.Pos0 != "助詞" {
		t.Fatal("Pos0 error")
	}
	if len(m.Features) != 3 {
		t.Errorf("Expected the number of features is 3 but got %v", m.Features)
	} else if _, ok := m.Features["かな漢字"]; !ok {
		t.Errorf("Feautre かな漢字 not found")
	} else if _, ok := m.Features["ひらがな"]; !ok {
		t.Errorf("Feautre ひらがな not found")
	} else if _, ok := m.Features["付属"]; !ok {
		t.Errorf("Feautre 付属 not found")
	}
}

func TestMorphemeKNP(t *testing.T) {
	line := "構文 こうぶん 構文 名詞 6 普通名詞 1 * 0 * 0 \"代表表記:構文/こうぶん カテゴリ:抽象物\" " + sampleFeature
	m, err := NewMorpheme(line)

	if err != nil {
		t.Fatal(err)
	}

	gf := getFeatures(sampleFeature, '>', 1)
	if m.Surface != "構文" {
		t.Errorf("Midashi Error\n")
	} else if m.CFormID != 0 {
		t.Errorf("Katsuyou2ID Error\n")
	} else if m.Rep != "構文/こうぶん" {
		t.Errorf("Rep Error\n")
	} else if !reflect.DeepEqual(m.Features, gf) {
		t.Errorf("Features Error [%v] != [%v]\n", m.Features, gf)
	}
}
