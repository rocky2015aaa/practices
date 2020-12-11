package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(Solution("a0Ba"))
	fmt.Println(Solution("a00b"))
}

func Solution(S string) int {
	max := -1
	for i := range S {
		for j := i; j <= len(S); j++ {
			temp := S[i:j]
			match, _ := regexp.MatchString("^[a-z]*[A-Z]+[a-z]*$", temp)
			strLen := len(temp)
			if match && max < strLen {
				max = strLen
			}
		}
	}
	return max
}
