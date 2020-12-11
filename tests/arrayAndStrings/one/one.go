package one

func CheckAllUnique(str string) bool {
	for i := range str {
		var count int
		for j := i; j < len(str); j++ {
			if str[i] == str[j] {
				count++
			}

			if count > 1 {
				return false
			}
		}
	}

	return true
}
