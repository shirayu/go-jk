package jk

import (
	"errors"
	"strconv"
	"strings"
)

//Argument represent a single argument
type Argument struct {
	Sid string
	Tid int
	Rep string
}

// Arguments is a slice of Argument
type Arguments []Argument

//ArgMap is a map from string to Arguments
type ArgMap map[string]Arguments

// Pas represents a predicate argument structure
type Pas struct {
	Cfid string
	Args ArgMap
}

// NewPas returns a Pas
func NewPas(line string, knpstyle bool) (*Pas, error) {
	if !knpstyle {
		return nil, errors.New("Not Implemented")
	}
	return parseKnpStylePas(line)
}

func parseKnpStylePas(val string) (*Pas, error) {
	pas := new(Pas)
	pas.Args = ArgMap{}
	c0 := strings.Index(val, ":")
	c1 := c0 + 1 + strings.Index(val[c0+1:], ":")
	pas.Cfid = val[:c0] + ":" + val[c0+1:c1]

	if strings.Count(val, ":") < 2 {
		return nil, nil
	}

	ks := strings.Split(val[c1+1:], ";")
	for _, k := range ks {
		items := strings.Split(k, "/")
		casetype := items[1]
		if casetype == "U" || casetype == "-" {
			continue
		}
		if len(items) != 6 {
			return nil, errors.New("Invalid number of values")
		}

		mycase := items[0]
		rep := items[2]
		tid, err := strconv.Atoi(items[3])
		if err != nil {
			return nil, err
		}
		sid := items[5]
		arg := Argument{Sid: sid, Tid: tid, Rep: rep}

		as, ok := pas.Args[mycase]
		if !ok {
			as = Arguments{}
			pas.Args[mycase] = as
		}
		pas.Args[mycase] = append(pas.Args[mycase], arg)
	}
	return pas, nil
}
