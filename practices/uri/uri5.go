package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var inputString string
	fmt.Scanf("%s", &inputString)
	numOfInputs, _ := strconv.Atoi(inputString)
	inputStrings := []string{}
	for i := 0; i < numOfInputs; i++ {
		fmt.Scanf("%s", &inputString)
		inputStrings = append(inputStrings, inputString)
	}

	for index, _ := range inputStrings {
		for i := 1; i < 128; i++ {
			if i != 25 {
				char := fmt.Sprintf("%c", i)
				charHex := fmt.Sprintf("%x", char)
				if len(charHex) == 1 {
					charHex = "0" + charHex
				}
				charHex = "%" + charHex
				inputStrings[index] = strings.Replace(inputStrings[index], charHex, char, -1)
			}
		}
		inputStrings[index] = strings.Replace(inputStrings[index], "%25", "%", 0)
	}

	for _, inputStrings := range inputStrings {
		fmt.Printf("%s\n", inputStrings)
	}
}
