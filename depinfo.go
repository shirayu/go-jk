package jk

import (
	"errors"
	"strconv"
	"strings"
)

type DependencyInfo struct {
	To       int
	DepType  rune
	Features Features
}

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

	self := new(DependencyInfo)
	var err error
	self.To, err = strconv.Atoi(line[sep1+1 : sep2-1])
	if err != nil {
		return nil, err
	}
	self.DepType = rune(line[sep2-1])

	self.Features = GetFeatures(line[sep2+1:])

	return self, err
}

type DependencyInfos []*DependencyInfo
