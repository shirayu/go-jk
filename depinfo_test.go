package jk

import (
	"reflect"
	"testing"
)

func TestDependencyInfo(t *testing.T) {
	f := `<BGH:パン/ぱん><文頭><ガ><助詞><体言><係:ガ格><区切:0-0><格要素><連用要素><正規化代表表記:パン/ぱん><主辞代表表記:パン/ぱん>`
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

}
