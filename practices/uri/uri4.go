package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
 - uri encoding converter
   The program to convert encoded characters to special characters with uri
 #input : number of uris, uris each
 #output : coverted uris with special characters
*/

func main() {
	input := func() string {
		var inputString string
		fmt.Scanf("%s", &inputString)
		return inputString
	}

	numOfInputs, _ := strconv.Atoi(input())
	inputStrings := make([]string, numOfInputs, numOfInputs)
	for i := 0; i < numOfInputs; i++ {
		inputString := input()
		inputStrings = append(inputStrings, inputString)
	}

	newUris := ConvertUri(inputStrings[1:])
	for _, newUri := range newUris {
		fmt.Printf("%s\n", newUri)
	}
}

func ConvertUri(inputStrings []string) []string {
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

	return inputStrings
}
