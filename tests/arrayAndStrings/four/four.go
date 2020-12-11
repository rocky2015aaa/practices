package four

func EncodingSpace(str string) string {
	strArr := []rune(str)
	var newStr string
	for i := 0; i < len(strArr); i++ {
		if strArr[i] != ' ' {
			if i-1 >= 0 && strArr[i-1] == ' ' {
				newStr += "%20"
			}
			newStr += string(strArr[i])
		}
	}

	return newStr
}
