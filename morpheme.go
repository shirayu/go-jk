package jk

import (
	"strconv"
	"strings"
)

//Morpheme is a morpheme
type Morpheme struct {
	Doukeis       Morphemes
	Surface       string   //Midashi
	Pronunciation string   //Yomi
	RootForm      string   //Genkei
	Pos0          string   //Hinshi
	Pos0ID        int      //HinshiID
	Pos1          string   //Bunrui
	Pos1ID        int      //BunruiID
	CType         string   //Katsuyou1
	CTypeID       int      //Katsuyou1ID
	CForm         string   //Katsuyou2
	CFormID       int      //Katsuyou2ID
	Seminfo       string   //Seminfo
	Rep           string   //Rep
	Features      Features //Features by KNP
}

//NewMorpheme returns a morpheme for the given line
func NewMorpheme(line string) (*Morpheme, error) {
	mrph := new(Morpheme)
	if strings.HasPrefix(line, " ") { //for JUMAN blank
		line = "\\" + line
	}

	seps := []int{}
	var prevC rune
	for i, c := range line {
		if c == ' ' && prevC != '\\' {
			seps = append(seps, i)
			if len(seps) == 3 {
				break
			}
		}
		prevC = c
	}

	mrph.Surface = strings.Replace(line[0:seps[0]], "\\ ", " ", -1)
	mrph.Pronunciation = strings.Replace(line[seps[0]+1:seps[1]], "\\ ", " ", -1)
	mrph.RootForm = strings.Replace(line[seps[1]+1:seps[2]], "\\ ", " ", -1)

	items := strings.SplitN(line[seps[2]+1:], " ", 9)
	mrph.Pos0 = items[0]

	hinshiID, err := strconv.Atoi(items[1])
	if err != nil {
		return nil, err
	}
	mrph.Pos0ID = hinshiID

	mrph.Pos1 = items[2]
	bunruiID, err := strconv.Atoi(items[3])
	if err != nil {
		return nil, err
	}
	mrph.Pos1ID = bunruiID

	mrph.CType = items[4]
	katsuyo1ID, err := strconv.Atoi(items[5])
	if err != nil {
		return nil, err
	}
	mrph.CTypeID = katsuyo1ID

	mrph.CForm = items[6]
	katsuyo2ID, err := strconv.Atoi(items[7])
	if err != nil {
		return nil, err
	}
	mrph.CFormID = katsuyo2ID

	rest := items[8]
	seminfoStartPos := strings.Index(rest, "\"")
	featureStart := 0
	if seminfoStartPos == -1 {
		mrph.Seminfo = ""
		mrph.Rep = mrph.RootForm + "/" + mrph.RootForm
		featureStart = strings.Index(rest, "<")
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
			mrph.Rep = mrph.RootForm + "/" + mrph.RootForm
		}
		featureStart = seminfoStartPos + 1 + seminfoCharNum + 2
	}

	if featureStart >= 0 && featureStart < len(rest) {
		mrph.Features = getFeatures(rest[featureStart:], '>', 1)
	}

	return mrph, err
}
