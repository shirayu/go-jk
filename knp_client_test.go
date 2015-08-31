package jk

import (
	"strings"
	"testing"
)

func TestKnp(t *testing.T) {
	knp, err := NewKnpClient("localhost:31000")
	if err != nil {
		t.Fatal("Error to open the knp socket: ", err)
	}

	input := `パン ぱん パン 名詞 6 普通名詞 1 * 0 * 0 "代表表記:パン/ぱん カテゴリ:人工物-食べ物 ドメイン:料理・食事"
が が が 助詞 9 格助詞 1 * 0 * 0 NIL
食べ たべ 食べる 動詞 2 * 0 母音動詞 1 未然形 3 "代表表記:食べる/たべる ドメイン:料理・食事"
られる られる られる 接尾辞 14 動詞性接尾辞 7 母音動詞 1 基本形 2 "代表表記:られる/られる"
EOS`

	ret_lines, err := knp.RawParse(input)
	if err != nil {
		t.Error("Error to parse [%v]", err)
	}
	if c := strings.Count(ret_lines, "\n"); c != 9 {
		t.Errorf("expceted length is 9 but %d", c)
	}

	s, err := knp.Parse(input)
	if err != nil {
		t.Error("Error to parse [%v]", err)
	}
	if s.Len() != 4 {
		t.Errorf("expceted length is 4 but %d", s.Len())
	}

}
