package config

import (
	"fmt"
	"regexp"
	"strings"
)

type Smiles struct {
	list  []string
	regex *regexp.Regexp
}

func (sm *Smiles) RemoveFromString(s string) string {
	return sm.regex.ReplaceAllLiteralString(s, "")
}

func MakeSmiles(list []string) *Smiles {
	smiles := &Smiles{
		list:  list,
		regex: regexp.MustCompile(fmt.Sprintf(":(%s):", strings.Join(list, "|"))),
	}
	return smiles
}
