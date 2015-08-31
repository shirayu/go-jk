package jk

import (
	"reflect"
	"testing"
)

func TestSentence(t *testing.T) {
	lines_string := `みて みて みる 動詞 2 * 0 母音動詞 1 タ系連用テ形 14 "代表表記:見る/みる 補文ト 自他動詞:自:見える/みえる"
@ みて みて みる 動詞 2 * 0 母音動詞 1 タ系連用テ形 14 "代表表記:診る/みる 補文ト ドメイン:健康・医学"
いる いる いる 接尾辞 14 動詞性接尾辞 7 母音動詞 1 基本形 2 "代表表記:いる/いる"`
	s, err := NewSentence(lines_string)
	if err != nil {
		t.Fatal(err)
	}

	if s.Len() != 2 {
		t.Errorf("expeceted length of morphemes is 2, but %d\n", s.Len())
	}

}

func TestKnpSentence(t *testing.T) {
	knp_lines := `# S-ID:2 KNP:4.2-17ed40d DATE:2015/08/28 SCORE:-8.99590
* 1D <BGH:パン/ぱん><文頭><ガ><助詞><体言><係:ガ格><区切:0-0><格要素><連用要素><正規化代表表記:パン/ぱん><主辞代表表記:パン/ぱん>
+ 1D <BGH:パン/ぱん><文頭><ガ><助詞><体言><係:ガ格><区切:0-0><格要素><連用要素><名詞項候補><先行詞候補><正規化代表表記:パン/ぱん><解析格:ガ>
パン ぱん パン 名詞 6 普通名詞 1 * 0 * 0 "代表表記:パン/ぱん カテゴリ:人工物-食べ物 ドメイン:料理・食事" <代表表記:パン/ぱん><カテゴリ:人工物-食べ物><ドメイン:料理・食事><正規化代表表記:パン/ぱん><記英数カ><カタカナ><名詞相当語><文頭><自立><内容語><タグ単位始><文節始><固有キー><文節主辞>
が が が 助詞 9 格助詞 1 * 0 * 0 NIL <かな漢字><ひらがな><付属>
* -1D <BGH:食べる/たべる><文末><態:受動|可能><〜られる><用言:動><レベル:C><区切:5-5><ID:（文末）><提題受:30><主節><動態述語><正規化代表表記:食べる/たべる><主辞代表表記:食べる/たべる>
+ -1D <BGH:食べる/たべる><文末><態:受動|可能><〜られる><用言:動><レベル:C><区切:5-5><ID:（文末）><提題受:30><主節><動態述語><正規化代表表記:食べる/たべる><用言代表表記:食べる/たべる+られる/られる><時制-未来><主題格:一人称優位><格関係0:ガ:パン><格解析結果:食べる/たべる+られる/られる:動1:ガ/C/パン/0/0/2;ニ/U/-/-/-/-;デ/U/-/-/-/-;カラ/U/-/-/-/-;時間/U/-/-/-/-;ノ/U/-/-/-/-;ガ２/U/-/-/-/->
食べ たべ 食べる 動詞 2 * 0 母音動詞 1 未然形 3 "代表表記:食べる/たべる ドメイン:料理・食事" <代表表記:食べる/たべる><ドメイン:料理・食事><正規化代表表記:食べる/たべる><かな漢字><活用語><自立><内容語><タグ単位始><文節始><文節主辞>
られる られる られる 接尾辞 14 動詞性接尾辞 7 母音動詞 1 基本形 2 "代表表記:られる/られる" <代表表記:られる/られる><正規化代表表記:られる/られる><かな漢字><ひらがな><活用語><文末><表現文末><付属>
EOS`
	gold_morphemes := Morphemes{
		&Morpheme{Midashi: "パン", Yomi: "ぱん", Genkei: "パン",
			Hinshi: "名詞", Bunrui: "普通名詞", Katsuyou1: "*", Katsuyou2: "*",
			Seminfo:  "代表表記:パン/ぱん カテゴリ:人工物-食べ物 ドメイン:料理・食事",
			Rep:      "パン/ぱん",
			Features: GetFeatures(`<代表表記:パン/ぱん><カテゴリ:人工物-食べ物><ドメイン:料理・食事><正規化代表表記:パン/ぱん><記英数カ><カタカナ><名詞相当語><文頭><自立><内容語><タグ単位始><文節始><固有キー><文節主辞>`)},
		&Morpheme{Midashi: "が", Yomi: "が", Genkei: "が",
			Hinshi: "助詞", Bunrui: "格助詞", Katsuyou1: "*", Katsuyou2: "*",
			Rep: "が/が", Features: GetFeatures(``),
		},
		&Morpheme{Midashi: "食べ", Yomi: "たべ", Genkei: "食べる",
			Hinshi: "動詞", Bunrui: "*", Katsuyou1: "母音動詞", Katsuyou2: "未然形",
			Seminfo:  `代表表記:食べる/たべる ドメイン:料理・食事`,
			Rep:      "食べる/たべる",
			Features: GetFeatures(`<代表表記:食べる/たべる><ドメイン:料理・食事><正規化代表表記:食べる/たべる><かな漢字><活用語><自立><内容語><タグ単位始><文節始><文節主辞>`),
		},
		&Morpheme{Midashi: "られる", Yomi: "られる", Genkei: "られる",
			Hinshi: "接尾辞", Bunrui: "動詞性接尾辞", Katsuyou1: "母音動詞", Katsuyou2: "基本形",
			Seminfo:  `代表表記:られる/られる`,
			Rep:      `られる/られる`,
			Features: GetFeatures(`<代表表記:られる/られる><正規化代表表記:られる/られる><かな漢字><ひらがな><活用語><文末><表現文末><付属>`)},
	}

	s, err := NewSentence(knp_lines)
	if err != nil {
		t.Fatal(err)
	}

	if s.Len() != 4 {
		t.Errorf("expeceted length of morphemes is 4, but %d\n", s.Len())
	}

	for i, sys := range s.Morphemes {
		gold := gold_morphemes[i]
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

}
