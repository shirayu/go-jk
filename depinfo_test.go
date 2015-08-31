package jk

import (
	"reflect"
	"testing"
)

func TestDependencyInfo(t *testing.T) {
	f := `+ -1D <BGH:食べる/たべる><文末><態:受動|可能><〜られる><用言:動><レベル:C><区切:5-5><ID:（文末）><提題受:30><主節><動態述語><正規化代表表記:食べる/たべる><用言代表表記:食べる/たべる+られる/られる><時制-未来><主題格:一人称優位><格関係0:ガ:パン><格解析結果:食べる/たべる+られる/られる:動1:ガ/C/パン/0/0/2;ニ/U/-/-/-/-;デ/U/-/-/-/-;カラ/U/-/-/-/-;時間/U/-/-/-/-;ノ/U/-/-/-/-;ガ２/U/-/-/-/-><BGH:パン/ぱん><文頭><ガ><助詞><体言><係:ガ格><区切:0-0><格要素><連用要素><正規化代表表記:パン/ぱん><主辞代表表記:パン/ぱん>`
	line := `* 1D ` + f

	dp, err := NewDependencyInfo(line)
	if err != nil {
		t.Errorf("Got error: %v", err)
		return
	}
	if dp.To != 1 {
		t.Errorf("To error")
	}
	if dp.DepType != 'D' {
		t.Errorf("DepType error")
	}
	if !reflect.DeepEqual(dp.Features, GetFeatures(f)) {
		t.Errorf("Features error")
	}

	if dp.GetPredRep() != "食べる/たべる+られる/られる:動" {
		t.Errorf("PredRep error")
	}

}
