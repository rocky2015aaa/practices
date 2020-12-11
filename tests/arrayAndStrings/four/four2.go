package four

import "strings"

func EncodingSpace2(str string) string {
	str = strings.TrimSpace(str)
	strLeng := len(str)
	var spaceCount int
	for i := 0; i < strLeng; i++ {
		if str[i] == ' ' {
			spaceCount++
		}
	}

	newCapacity := strLeng + spaceCount*2
	newStr := make([]byte, newCapacity, newCapacity)

	for j := strLeng - 1; j >= 0; j-- {
		if str[j] == ' ' {
			newStr[newCapacity-1] = '0'
			newStr[newCapacity-2] = '2'
			newStr[newCapacity-3] = '%'
			newCapacity -= 3
		} else {
			newStr[newCapacity-1] = str[j]
			newCapacity -= 1
		}
	}

	return string(newStr)
}
