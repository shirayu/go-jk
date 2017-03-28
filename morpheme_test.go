package jk

import (
	"reflect"
	"testing"
)

func checkMrph(t *testing.T, m, gold *Morpheme) {
	if m.Surface != gold.Surface {
		t.Errorf("Surface Error: expected %s but got %s", gold.Surface, m.Surface)
	}
	if m.Pronunciation != gold.Pronunciation {
		t.Errorf("Pronunciation Error: expected %s but got %s", gold.Pronunciation, m.Pronunciation)
	}
	if m.RootForm != gold.RootForm {
		t.Errorf("RootForm Error: expected %s but got %s", gold.RootForm, m.RootForm)
	}
	if m.Pos0 != gold.Pos0 {
		t.Errorf("Pos0 Error: expected %s but got %s", gold.Pos0, m.Pos0)
	}
	if m.Pos0ID != gold.Pos0ID {
		t.Errorf("Pos0ID Error: expected %d but got %d", gold.Pos0ID, m.Pos0ID)
	}
	if m.Pos1 != gold.Pos1 {
		t.Errorf("Pos1 Error: expected %s but got %s", gold.Pos1, m.Pos1)
	}
	if m.Pos1ID != gold.Pos1ID {
		t.Errorf("Pos1ID Error: expected %d but got %d", gold.Pos1ID, m.Pos1ID)
	}
	if m.CType != gold.CType {
		t.Errorf("CType Error: expected %s but got %s", gold.CType, m.CType)
	}
	if m.CTypeID != gold.CTypeID {
		t.Errorf("CTypeID Error: expected %d but got %d", gold.CTypeID, m.CTypeID)
	}
	if m.CForm != gold.CForm {
		t.Errorf("CForm Error: expected %s but got %s", gold.CForm, m.CForm)
	}
	if m.CFormID != gold.CFormID {
		t.Errorf("CFormID Error: expected %d but got %d", gold.CFormID, m.CFormID)
	}
	if m.Seminfo != gold.Seminfo {
		t.Errorf("Seminfo Error: expected %s but got %s", gold.Seminfo, m.Seminfo)
	}
	if m.Rep != gold.Rep {
		t.Errorf("Rep Error: expected %s but got %s", gold.Rep, m.Rep)
	}
	if !reflect.DeepEqual(m.Doukeis, gold.Doukeis) {
		t.Errorf("Doukeis Error: expected %v but got %v", gold.Doukeis, m.Doukeis)
	}
	if !reflect.DeepEqual(m.Features, gold.Features) {
		t.Errorf("Features Error: expected %v but got %v", gold.Features, m.Features)
	}
}

func TestMorpheme(t *testing.T) {
	tests := []struct {
		line string
		gold Morpheme
	}{
		{
			line: "探して さがして 探す 動詞 2 * 0 子音動詞サ行 5 タ系連用テ形 14 \"代表表記:探す/さがす\"",
			gold: Morpheme{
				Surface:       "探して",
				Pronunciation: "さがして",
				RootForm:      "探す",
				Pos0:          "動詞",
				Pos0ID:        2,
				Pos1:          "*",
				Pos1ID:        0,
				CType:         "子音動詞サ行",
				CTypeID:       5,
				CForm:         "タ系連用テ形",
				CFormID:       14,
				Seminfo:       "代表表記:探す/さがす",
				Rep:           "探す/さがす",
			},
		},
		{
			line: `を を を 助詞 9 格助詞 1 * 0 * 0 NIL <かな漢字><ひらがな><付属>`,
			gold: Morpheme{
				Surface:       "を",
				Pronunciation: "を",
				RootForm:      "を",
				Pos0:          "助詞",
				Pos0ID:        9,
				Pos1:          "格助詞",
				Pos1ID:        1,
				CType:         "*",
				CTypeID:       0,
				CForm:         "*",
				CFormID:       0,
				Seminfo:       "",
				Rep:           "を/を",
				Features:      Features{"かな漢字": "", "ひらがな": "", "付属": ""},
			},
		},
		{ // KNP style
			line: "構文 こうぶん 構文 名詞 6 普通名詞 1 * 0 * 0 \"代表表記:構文/こうぶん カテゴリ:抽象物\" " + sampleFeature,
			gold: Morpheme{
				Surface:       "構文",
				Pronunciation: "こうぶん",
				RootForm:      "構文",
				Pos0:          "名詞",
				Pos0ID:        6,
				Pos1:          "普通名詞",
				Pos1ID:        1,
				CType:         "*",
				CTypeID:       0,
				CForm:         "*",
				CFormID:       0,
				Seminfo:       "代表表記:構文/こうぶん カテゴリ:抽象物",
				Rep:           "構文/こうぶん",
				Features:      Features{`代表表記`: `構文/こうぶん`, `カテゴリ`: `抽象物`, `正規化代表表記`: `構文/こうぶん`, `漢字`: ``},
			},
		},
		{ // blank in JUMAN
			line: `  \  \  特殊 1 空白 6 * 0 * 0 NIL`,
			gold: Morpheme{
				Surface:       " ",
				Pronunciation: " ",
				RootForm:      " ",
				Pos0:          "特殊",
				Pos0ID:        1,
				Pos1:          "空白",
				Pos1ID:        6,
				CType:         "*",
				CTypeID:       0,
				CForm:         "*",
				CFormID:       0,
				Rep:           " / ",
			},
		},
		{ // blank in JUMAN++
			line: `\  \  \  特殊 1 空白 6 * 0 * 0 "代表表記: / "`,
			gold: Morpheme{
				Surface:       " ",
				Pronunciation: " ",
				RootForm:      " ",
				Pos0:          "特殊",
				Pos0ID:        1,
				Pos1:          "空白",
				Pos1ID:        6,
				CType:         "*",
				CTypeID:       0,
				CForm:         "*",
				CFormID:       0,
				Seminfo:       "代表表記: / ",
				Rep:           "",
			},
		},
		{ // morpheme with blank in JUMAN++
			line: `\ X \ X \ X 未定義語 15 その他 1 * 0 * 0 "品詞推定:名詞"`,
			gold: Morpheme{
				Surface:       " X",
				Pronunciation: " X",
				RootForm:      " X",
				Pos0:          "未定義語",
				Pos0ID:        15,
				Pos1:          "その他",
				Pos1ID:        1,
				CType:         "*",
				CTypeID:       0,
				CForm:         "*",
				CFormID:       0,
				Rep:           " X/ X",
				Seminfo:       `品詞推定:名詞`,
				//                 Features:      Features{`品詞推定`: `名詞`}, //TODO fix me
			},
		},
	}

	for _, test := range tests {
		m, err := NewMorpheme(test.line)
		if err != nil {
			t.Fatal(err)
		}
		checkMrph(t, m, &test.gold)
	}
}
