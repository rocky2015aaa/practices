package two

func Reverse2(str string) string {
	if len(str) == 0 {
		return str
	}

	return Reverse2(str[1:]) + string(str[0])
}
