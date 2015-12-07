package jk

import (
	"errors"
	"strconv"
	"strings"
)

//DependencyInfo handles dependency information of bunsetsu or basic phrases
type DependencyInfo struct {
	To       int
	DepType  rune
	Features Features
}

//NewDependencyInfo creates a new DependencyInfo with the given line
func NewDependencyInfo(line string) (*DependencyInfo, error) {
	sep1 := strings.IndexRune(line, ' ')
	if sep1 != 1 {
		return nil, errors.New("Invalid separator")
	}
	sep2 := strings.IndexRune(line[sep1+1:], ' ')
	if sep2 < 0 {
		return nil, errors.New("Invalid separator")
	}
	sep2 += sep1 + 1

	depi := new(DependencyInfo)
	var err error
	depi.To, err = strconv.Atoi(line[sep1+1 : sep2-1])
	if err != nil {
		return nil, err
	}
	depi.DepType = rune(line[sep2-1])

	depi.Features = GetFeatures(line[sep2+1:])

	return depi, err
}

//DependencyInfos is a slice of DependencyInfo
type DependencyInfos []*DependencyInfo

//GetPredRep returns the "rep" for the DependencyInfo
func (depi *DependencyInfo) GetPredRep() string {
	pname, ok := depi.Features["用言代表表記"]
	if !ok {
		return ""
	}
	vtype, ok := depi.Features["用言"]
	if !ok {
		return ""
	}

	return pname + ":" + vtype
}
