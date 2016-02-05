package jk

import (
	"reflect"
	"strings"
	"testing"
)

func TestSentence(t *testing.T) {
	linesString := `みて みて みる 動詞 2 * 0 母音動詞 1 タ系連用テ形 14 "代表表記:見る/みる 補文ト 自他動詞:自:見える/みえる"
@ みて みて みる 動詞 2 * 0 母音動詞 1 タ系連用テ形 14 "代表表記:診る/みる 補文ト ドメイン:健康・医学"
いる いる いる 接尾辞 14 動詞性接尾辞 7 母音動詞 1 基本形 2 "代表表記:いる/いる"`
	s, err := NewSentence(strings.Split(linesString, "\n"))
	if err != nil {
		t.Fatal(err)
	}

	if s.Len() != 2 {
		t.Errorf("expeceted length of morphemes is 2, but %d\n", s.Len())
	}

}

func TestKnpSentence(t *testing.T) {
	goldMorphemes := Morphemes{
		&Morpheme{Surface: "パン", Pronunciation: "ぱん", RootForm: "パン",
			Pos0: "名詞", Pos1: "普通名詞", CType: "*", CForm: "*",
			Seminfo:  "代表表記:パン/ぱん カテゴリ:人工物-食べ物 ドメイン:料理・食事",
			Rep:      "パン/ぱん",
			Features: getFeatures(`<代表表記:パン/ぱん><カテゴリ:人工物-食べ物><ドメイン:料理・食事><正規化代表表記:パン/ぱん><記英数カ><カタカナ><名詞相当語><文頭><自立><内容語><タグ単位始><文節始><固有キー><文節主辞>`, '>', 1)},
		&Morpheme{Surface: "が", Pronunciation: "が", RootForm: "が",
			Pos0: "助詞", Pos1: "格助詞", CType: "*", CForm: "*",
			Rep: "が/が", Features: getFeatures(``, '>', 1),
		},
		&Morpheme{Surface: "食べ", Pronunciation: "たべ", RootForm: "食べる",
			Pos0: "動詞", Pos1: "*", CType: "母音動詞", CForm: "未然形",
			Seminfo:  `代表表記:食べる/たべる ドメイン:料理・食事`,
			Rep:      "食べる/たべる",
			Features: getFeatures(`<代表表記:食べる/たべる><ドメイン:料理・食事><正規化代表表記:食べる/たべる><かな漢字><活用語><自立><内容語><タグ単位始><文節始><文節主辞>`, '>', 1),
		},
		&Morpheme{Surface: "られる", Pronunciation: "られる", RootForm: "られる",
			Pos0: "接尾辞", Pos1: "動詞性接尾辞", CType: "母音動詞", CForm: "基本形",
			Seminfo:  `代表表記:られる/られる`,
			Rep:      `られる/られる`,
			Features: getFeatures(`<代表表記:られる/られる><正規化代表表記:られる/られる><かな漢字><ひらがな><活用語><文末><表現文末><付属>`, '>', 1)},
	}

	s, err := NewSentence(strings.Split(sampleKnpOutput, "\n"))
	if err != nil {
		t.Fatal(err)
	}

	if s.Len() != 4 {
		t.Errorf("expeceted length of morphemes is 4, but %d\n", s.Len())
	}

	if s.ID != sampleID {
		t.Errorf("expeceted ID is %s, but got %s", sampleID, s.ID)
	}

	for i, sys := range s.Morphemes {
		gold := goldMorphemes[i]
		if gold == nil {
			t.Errorf("gold error")
		}
		if sys == nil {
			t.Errorf("sys is nil")
		}
		if sys.Doukeis != nil {
			t.Errorf("Doukeis parse error")
		}

		sysvals := reflect.ValueOf(sys).Elem()
		goldvals := reflect.ValueOf(gold).Elem()
		num := goldvals.NumField()
		for i := 0; i < num; i++ {
			tagname := goldvals.Type().Field(i).Name
			sysv := sysvals.FieldByName(tagname).String()
			goldv := goldvals.FieldByName(tagname).String()
			if sysv != goldv {
				t.Errorf("%s parse error: [%v] expeceted but got [%v]", tagname, goldv, sysv)
			}
		}
	}

	if len(s.Bunsetsus) != 2 {
		t.Errorf("Bunsetsu size error: expeceted %v but got %v", len(s.Bunsetsus), 2)
	}
	if len(s.BasicPhrases) != 2 {
		t.Errorf("BasicPhrases size error: expeceted %v but got %v", len(s.BasicPhrases), 2)
	}
	if s.BasicPhrases[1].Features["用言代表表記"] != `食べる/たべる+られる/られる` {
		t.Errorf("Features error")
	}
	if s.BasicPhrases[1].Features["用言"] != `動` {
		t.Errorf("Features error")
	}

	if sys := s.MorphemePositions; !reflect.DeepEqual(sys, []int{0, 2, 3, 5, 8}) {
		t.Errorf("Got %v", sys)
	}
	if sys := s.BasicPhrasePositions; !reflect.DeepEqual(sys, []int{0, 3, 8}) {
		t.Errorf("Got %v", sys)
	}

	if sys := s.BasicPhraseMorphemeIndexs; !reflect.DeepEqual(sys, []int{0, 2, 4}) {
		t.Errorf("Got %v", sys)
	}
	if sys := s.GetMorphemes(-1); sys != nil {
		t.Errorf("Got %v", sys)
	} else if sys := s.GetMorphemes(0); sys == nil {
		t.Errorf("Got %v", sys)
	} else if sys := s.GetMorphemes(9999); sys != nil {
		t.Errorf("Got %v", sys)
	}

}
