package three

const NUM_OF_ASCIIS = 256

func CheckPermutationAscii(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	letters := make([]int, NUM_OF_ASCIIS, NUM_OF_ASCIIS)

	str1Array := []rune(str1)
	str2Array := []rune(str2)

	for _, str1Element := range str1Array {
		letters[str1Element]++
	}

	for i := 0; i < len(str2); i++ {
		index := str2Array[i]
		letters[index]--
		if letters[index] < 0 {
			return false
		}
	}

	return true
}
