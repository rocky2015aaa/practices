package main

import (
	"fmt"
)

func main() {
	inputNum := 2
	inputStrings := make([]string, inputNum, inputNum)

	/*
		var inputString string
		for i := 0; i < inputNum; i++ {
			fmt.Scanf("%s", &inputString)
			inputStrings = append(inputStrings, inputString)
		}
	*/

	for i := range inputStrings {
		fmt.Scanf("%s", &inputStrings[i])
	}

	for i, inputString := range inputStrings {
		fmt.Printf("%d %s\n", i, inputString)
	}
}
