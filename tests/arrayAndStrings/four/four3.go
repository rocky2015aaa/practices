package four

import "strings"

func EncodingSpace3(str string) string {
	str = strings.TrimSpace(str)
	strArr := []rune(str)
	var newStr string
	for i := 0; i < len(strArr); i++ {
		if strArr[i] == ' ' {
			newStr += "%20"
		} else {
			newStr += string(strArr[i])
		}
	}

	return newStr
}
