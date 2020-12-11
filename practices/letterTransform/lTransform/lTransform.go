package lTransform

import (
	"strings"
)

var CONVERT_MAP = map[string]string{
	"AB": "AA",
	"BA": "AA",
	"CB": "CC",
	"BC": "CC",
	"AA": "A",
	"CC": "C",
}

func Transform(S string) string {
	for {
		var hasSub bool
		for k, v := range CONVERT_MAP {
			index := strings.Index(S, k)
			if index > -1 {
				hasSub = true
				S = strings.Replace(S, k, v, 1)
			}
		}
		if !hasSub {
			break
		}
	}
	return S
}
