package three

func CheckPermutationAscii2(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	str1Arr := []rune(str1)
	str2Arr := []rune(str2)

	for i := 1; i < len(str1); i++ {
		str1Arr[0] += str1Arr[i]
		str2Arr[0] += str2Arr[i]
	}

	if str1Arr[0] != str2Arr[0] {
		return false
	}

	return true
}
