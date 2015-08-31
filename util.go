package jk

import (
	"strings"
)

func GetPredRep(ms Morphemes) string {
	ret := ""
	reps := []string{}
	vtype := "動"

	for idx, m := range ms {
		if len(ms) == 2 && idx == len(ms)-1 && m.Rep == "する/する" {
			break
		} else if idx == len(ms)-1 && m.Hinshi == "判定詞" {
			vtype = "判"
			break
		}
		reps = append(reps, m.Rep)

		if m.Hinshi == "形容詞" {
			vtype = "形"
		}
	}

	ret = strings.Join(reps, "+") + ":" + vtype

	return ret
}
