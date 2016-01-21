package jk

import (
	"strings"
)

//Features is a map of features
type Features map[string]string

//GetFeatures returns features for the given feature expression
func GetFeatures(line string, splitter rune, firstCharOffset int) Features {
	num := strings.Count(line, string(splitter))
	ret := make(Features, num)

	start := 0
	separator := 0 //0 means no separator found
	for i, char := range line {
		if char == splitter {
			k := line[start+firstCharOffset : i]
			v := ""
			if separator != 0 {
				k = line[start+firstCharOffset : separator]
				v = line[separator+1 : i]
			}
			ret[k] = v
			start = i + 1
			separator = 0
		} else if separator == 0 && char == ':' { //first separator
			separator = i
		}
	}

	//for last
	if firstCharOffset == 0 {
		k := line[start+firstCharOffset:]
		v := ""
		if separator != 0 {
			k = line[start+firstCharOffset : separator]
			v = line[separator+1:]
		}
		ret[k] = v
	}

	return ret
}
