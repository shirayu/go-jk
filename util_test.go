package jk

import (
	"testing"
)

type RepTestCase struct {
	Lines string
	Gold  string
}

func TestGetPredRep(t *testing.T) {

	cases := []RepTestCase{
		RepTestCase{`探す さがす 探す 動詞 2 * 0 子音動詞サ行 5 基本形 2 "代表表記:探す/さがす"`, "探す/さがす:動"},
		RepTestCase{`熱い あつい 熱い 形容詞 3 * 0 イ形容詞アウオ段 18 基本形 2 "代表表記:熱い/あつい 反義:形容詞:冷たい/つめたい"`, "熱い/あつい:形"},
		RepTestCase{`ググる ググる ググる 動詞 2 * 0 子音動詞ラ行 10 基本形 2 "自動獲得:テキスト"`, "ググる/ググる:動"},
		RepTestCase{`ググり ググり ググる 動詞 2 * 0 子音動詞ラ行 10 基本連用形 8 "自動獲得:テキスト"
たい たい たい 接尾辞 14 形容詞性述語接尾辞 5 イ形容詞アウオ段 18 基本形 2 "代表表記:たい/たい"`, "ググる/ググる+たい/たい:動"},
		RepTestCase{`判断 はんだん 判断 名詞 6 サ変名詞 2 * 0 * 0 "代表表記:判断/はんだん 補文ト カテゴリ:抽象物"`, "判断/はんだん:動"},
		RepTestCase{`判断 はんだん 判断 名詞 6 サ変名詞 2 * 0 * 0 "代表表記:判断/はんだん 補文ト カテゴリ:抽象物"
です です だ 判定詞 4 * 0 判定詞 25 デス列基本形 27 NIL`, "判断/はんだん:判"},
	}

	for _, testcase := range cases {
		s, err := NewSentence(testcase.Lines)
		if err != nil {
			t.Fatal(err)
		}

		rep := GetPredRep(s.GetMorphemes())
		if rep != testcase.Gold {
			t.Errorf("Expected %s but got %s", testcase.Gold, rep)
		}
	}

}
