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

	if m.Midashi != "探して" {
		t.Errorf("Midashi Error\n")
	} else if m.Katsuyou2ID != 14 {
		t.Errorf("Katsuyou2ID Error\n")
	} else if m.Seminfo != "代表表記:探す/さがす" {
		t.Errorf("Seminfo_id Error\n")
	} else if m.Rep != "探す/さがす" {
		t.Errorf("Rep Error\n")
	}
}

func TestMorphemeKNP(t *testing.T) {
	line := "構文 こうぶん 構文 名詞 6 普通名詞 1 * 0 * 0 \"代表表記:構文/こうぶん カテゴリ:抽象物\" " + sampleFeature
	m, err := NewMorpheme(line)

	if err != nil {
		t.Fatal(err)
	}

	gf := getFeatures(sampleFeature, '>', 1)
	if m.Midashi != "構文" {
		t.Errorf("Midashi Error\n")
	} else if m.Katsuyou2ID != 0 {
		t.Errorf("Katsuyou2ID Error\n")
	} else if m.Rep != "構文/こうぶん" {
		t.Errorf("Rep Error\n")
	} else if !reflect.DeepEqual(m.Features, gf) {
		t.Errorf("Features Error [%v] != [%v]\n", m.Features, gf)
	}
}
