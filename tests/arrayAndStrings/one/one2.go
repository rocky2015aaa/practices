package one

import (
	"sort"
	"strings"
)

func CheckAllUnique2(str string) bool {
	s := strings.Split(str, "")
	sort.Strings(s)
	previous := s[0]
	for i := 1; i < len(s); i++ {
		if previous == s[i] {
			return false
		} else {
			previous = s[i]
		}
	}

	return true
}
