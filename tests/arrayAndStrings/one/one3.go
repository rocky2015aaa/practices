package one

const (
	NUM_OF_ASCIIS = 256
)

func CheckAllUniqueAscii(str string) bool {
	strLen := len(str)
	if strLen > NUM_OF_ASCIIS {
		return false
	}

	charSet := make([]bool, NUM_OF_ASCIIS, NUM_OF_ASCIIS)
	for i := 0; i < strLen; i++ {
		val := int(str[i])
		if charSet[val] {
			return false
		}

		charSet[val] = true
	}

	return true
}
