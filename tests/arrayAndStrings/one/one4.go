package one

func CheckAllUniqueLowerCase(str string) bool {
	var checker int
	for i := 0; i < len(str); i++ {
		val := str[i] - 'a'
		if (checker & (1 << val)) > 0 {
			return false
		}
		checker |= (1 << val)
	}
	return true
}
