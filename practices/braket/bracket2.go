package main

import (
	"fmt"
)

var bracketPair = map[rune]rune{'{': '}', '[': ']', '(': ')'}

func main() {
	var numOfInput int
	fmt.Scanf("%d", &numOfInput)
	var input string
	brackets := []string{}
	for i := 0; i < numOfInput; i++ {
		fmt.Scanf("%s", &input)
		brackets = append(brackets, input)
	}
	for _, bracket := range brackets {
		fmt.Printf("%s\n", checkMatchBracket(bracket))
	}
}

func checkMatchBracket(bracket string) string {
	bracketStack := []rune{}
	for _, bracketEach := range bracket {
		top := len(bracketStack)
		if top > 0 && bracketEach == bracketPair[bracketStack[top-1]] {
			bracketStack = bracketStack[:top-1]
		} else {
			bracketStack = append(bracketStack, bracketEach)
		}
	}

	if len(bracketStack) > 0 {
		return "NO"
	}
	return "YES"
}
