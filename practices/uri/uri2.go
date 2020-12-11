package main

import (
	"fmt"
	"net/url"
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

	newUris := ConvertUri(inputStrings[1:])
	for _, newUri := range newUris {
		fmt.Printf("%s\n", newUri)
	}
}

func ConvertUri(inputStrings []string) []string {
	for index, _ := range inputStrings {
		u, _ := url.Parse(inputStrings[index])
		inputStrings[index] = strings.Replace(u.String(), u.RequestURI(), u.Path, -1)
	}
	return inputStrings
}
