package two

func Reverse(str string) string {
	strLeng := len(str)
	tempStr := []rune(str)
	for i := 0; i <= strLeng/2; i++ {
		tempStr[i], tempStr[strLeng-1-i] = tempStr[strLeng-1-i], tempStr[i]
	}

	return string(tempStr)
}
