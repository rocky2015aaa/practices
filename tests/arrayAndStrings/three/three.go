package three

import (
	"sort"
	"strings"
)

func CheckPermutation(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	char1 := strings.Split(str1, "")
	char2 := strings.Split(str2, "")
	sort.Strings(char1)
	sort.Strings(char2)
	for i := range char1 {
		if char1[i] != char2[i] {
			return false
		}
	}

	return true
}
