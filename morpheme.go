package jk

import (
	"strconv"
	"strings"
)

type Features map[string]string

type Morpheme struct {
	Doukeis      Morphemes
	Midashi      string
	Yomi         string
	Genkei       string
	Hinshi       string
	Hinshi_id    int
	Bunrui       string
	Bunrui_id    int
	Katsuyou1    string
	Katsuyou1_id int
	Katsuyou2    string
	Katsuyou2_id int
	Seminfo      string
	Rep          string
	Features     Features
}

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

func NewMorpheme(line string) (*Morpheme, error) {
	self := new(Morpheme)
	items := strings.SplitN(line, " ", 12)

	self.Midashi = items[0]
	self.Yomi = items[1]
	self.Genkei = items[2]
	self.Hinshi = items[3]

	hinshi_id, err := strconv.Atoi(items[4])
	if err != nil {
		return nil, err
	}
	self.Hinshi_id = hinshi_id

	self.Bunrui = items[5]
	bunrui_id, err := strconv.Atoi(items[6])
	if err != nil {
		return nil, err
	}
	self.Bunrui_id = bunrui_id

	self.Katsuyou1 = items[7]
	katsuyo1_id, err := strconv.Atoi(items[8])
	if err != nil {
		return nil, err
	}
	self.Katsuyou1_id = katsuyo1_id

	self.Katsuyou2 = items[9]
	katsuyo2_id, err := strconv.Atoi(items[10])
	if err != nil {
		return nil, err
	}
	self.Katsuyou2_id = katsuyo2_id

	rest := items[11]
	seminfo_start_pos := strings.Index(rest, "\"")
	if seminfo_start_pos == -1 {
		self.Seminfo = ""
		self.Rep = self.Genkei + "/" + self.Genkei
	} else {
		seminfo_char_num := strings.Index(rest[seminfo_start_pos+1:], "\"")
		self.Seminfo = rest[seminfo_start_pos+1 : seminfo_start_pos+1+seminfo_char_num]

		ret_name := "代表表記:"
		rep_start := strings.Index(self.Seminfo, ret_name)
		if rep_start != -1 {
			rep_end := strings.Index(self.Seminfo[rep_start:], " ")
			if rep_end == -1 {
				rep_end = len(self.Seminfo)
			} else {
				rep_end += rep_start
			}
			self.Rep = self.Seminfo[rep_start+len(ret_name) : rep_end]
		} else {
			self.Rep = self.Genkei + "/" + self.Genkei
		}

		feature_start := seminfo_start_pos + 1 + seminfo_char_num + 2
		if feature_start < len(rest) {
			self.Features = GetFeatures(rest[feature_start:])
		}

	}

	return self, err
}
