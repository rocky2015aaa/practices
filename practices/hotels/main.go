package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"practices/hotels/hotels"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	keywords := strings.Fields(input)

	var numOfInput int
	fmt.Scanf("%d", &numOfInput)

	reviews := make([]string, numOfInput, numOfInput)
	hotelIds := make([]int, numOfInput, numOfInput)
	for i := 0; i < numOfInput; i++ {
		fmt.Scanf("%d", &hotelIds[i])
		reviews[i], _ = reader.ReadString('\n')
	}

	fmt.Println(hotels.SortHotels(keywords, reviews, hotelIds))
}
