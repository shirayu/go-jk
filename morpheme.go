package jk

import (
	"strconv"
	"strings"
)

//Morpheme is a morpheme
type Morpheme struct {
	Doukeis     Morphemes
	Midashi     string
	Yomi        string
	Genkei      string
	Hinshi      string
	HinshiID    int
	Bunrui      string
	BunruiID    int
	Katsuyou1   string
	Katsuyou1ID int
	Katsuyou2   string
	Katsuyou2ID int
	Seminfo     string
	Rep         string
	Features    Features
}

//NewMorpheme returns a morpheme for the given line
func NewMorpheme(line string) (*Morpheme, error) {
	mrph := new(Morpheme)
	items := strings.SplitN(line, " ", 12)

	mrph.Midashi = items[0]
	mrph.Yomi = items[1]
	mrph.Genkei = items[2]
	mrph.Hinshi = items[3]

	hinshiID, err := strconv.Atoi(items[4])
	if err != nil {
		return nil, err
	}
	mrph.HinshiID = hinshiID

	mrph.Bunrui = items[5]
	bunruiID, err := strconv.Atoi(items[6])
	if err != nil {
		return nil, err
	}
	mrph.BunruiID = bunruiID

	mrph.Katsuyou1 = items[7]
	katsuyo1ID, err := strconv.Atoi(items[8])
	if err != nil {
		return nil, err
	}
	mrph.Katsuyou1ID = katsuyo1ID

	mrph.Katsuyou2 = items[9]
	katsuyo2ID, err := strconv.Atoi(items[10])
	if err != nil {
		return nil, err
	}
	mrph.Katsuyou2ID = katsuyo2ID

	rest := items[11]
	seminfoStartPos := strings.Index(rest, "\"")
	if seminfoStartPos == -1 {
		mrph.Seminfo = ""
		mrph.Rep = mrph.Genkei + "/" + mrph.Genkei
	} else {
		seminfoCharNum := strings.Index(rest[seminfoStartPos+1:], "\"")
		mrph.Seminfo = rest[seminfoStartPos+1 : seminfoStartPos+1+seminfoCharNum]

		retName := "代表表記:"
		repStart := strings.Index(mrph.Seminfo, retName)
		if repStart != -1 {
			repEnd := strings.Index(mrph.Seminfo[repStart:], " ")
			if repEnd == -1 {
				repEnd = len(mrph.Seminfo)
			} else {
				repEnd += repStart
			}
			mrph.Rep = mrph.Seminfo[repStart+len(retName) : repEnd]
		} else {
			mrph.Rep = mrph.Genkei + "/" + mrph.Genkei
		}

		featureStart := seminfoStartPos + 1 + seminfoCharNum + 2
		if featureStart < len(rest) {
			mrph.Features = GetFeatures(rest[featureStart:])
		}

	}

	return mrph, err
}
