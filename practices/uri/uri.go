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

const (
	INPUT_MIN = 1
	INPUT_MAX = 100
	URI_MAX   = 80
)

func main() {
	encodingMap := map[string]string{
		"%20": " ",
		"%21": "!",
		"%24": "$",
		"%25": "EOFEOFEOF",
		"%28": "(",
		"%29": ")",
		"%2a": "*",
	}

	input := func() string {
		var inputString string
		fmt.Scanf("%s", &inputString)
		return inputString
	}

	numOfInputs, err := strconv.Atoi(input())
	for err != nil || numOfInputs < INPUT_MIN || numOfInputs > INPUT_MAX {
		fmt.Printf("%s\n", "Invalid input. Try again.")
		numOfInputs, err = strconv.Atoi(input())
	}

	inputStrings := make([]string, numOfInputs, numOfInputs)
	for i := 0; i < numOfInputs; i++ {
		inputString := input()
		for len(inputString) > URI_MAX {
			fmt.Println("%s\n", "Out of length. Try again.")
			inputString = input()
		}
		inputStrings = append(inputStrings, inputString)
	}

	newUris := ConvertUri(inputStrings[1:], encodingMap)
	for _, newUri := range newUris {
		fmt.Printf("%s\n", newUri)
	}
}

func ConvertUri(inputStrings []string, encodingMap map[string]string) []string {
	for index, _ := range inputStrings {
		for encodingKey, encodingVal := range encodingMap {
			inputStrings[index] = strings.Replace(inputStrings[index], encodingKey, encodingVal, -1)
		}
		inputStrings[index] = strings.Replace(inputStrings[index], "EOFEOFEOF", "%", -1)
	}
	return inputStrings
}
