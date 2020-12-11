package solution

import (
	"strconv"
)

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A int, B int) int {
	// write your code in Go 1.4
	strA := strconv.Itoa(A)
	strB := strconv.Itoa(B)
	resultNum, err := strconv.Atoi(buildNumber(strA, strB))
	if err != nil || resultNum > 100000000 {
		return -1
	}

	return resultNum
}

func buildNumber(a, b string) string {
	resultNum := []rune{}
	lenA := len(a)
	lenB := len(b)
	var toStr string
	if lenA > lenB {
		toStr = a
	} else {
		toStr = b
	}

	for i := range toStr {
		if i < lenA {
			resultNum = append(resultNum, rune(a[i]))
		}
		if i < lenB {
			resultNum = append(resultNum, rune(b[i]))
		}
	}

	return string(resultNum)
}
