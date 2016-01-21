package jk

import (
	"strings"
)

//Features is a map of features
type Features map[string]string

//GetFeatures returns features for the given feature expression
func GetFeatures(line string) Features {
	num := strings.Count(line, "<")
	ret := make(Features, num)

	start := 0
	separator := 0
	for i, char := range line {
		if char == '>' {
			k := line[start+1 : i]
			v := ""
			if separator != 0 {
				k = line[start+1 : separator]
				v = line[separator+1 : i]
			}
			ret[k] = v
			start = i + 1
			separator = 0
		} else if separator == 0 && char == ':' { //first separator
			separator = i
		}
	}

	return ret
}
