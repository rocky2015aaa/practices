package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

var numbers []int

func init() {
	numbers = make([]int, 0, 0)
}

func main() {
	http.HandleFunc("/calculation", Calculation)
	http.HandleFunc("/price", StorePrice)
	if err := http.ListenAndServe(":9003", nil); err != nil {
		log.Fatal("failed to start server", err)
	}
}

func Calculation(resp http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		log.Printf("%s\n", err)
	}
	number, err := strconv.Atoi(request.PostFormValue("price"))
	if err != nil {
		log.Printf("%s\n", err)
	}

	resp.Header().Add("Access-Control-Allow-Origin", "http://localhost:9001")
	resp.Write([]byte(strconv.FormatFloat(calculator(numbers, number), 'f', 2, 32)))
}

func calculator(numbers []int, number int) float64 {
	lastIndex := len(numbers) - 1

	var result float64
	for i := lastIndex; i > lastIndex-number; i-- {
		result += float64(numbers[i])
	}
	average := (result / float64(number))

	log.Printf("total: %.2f, average: %.2f, number of input: %d\n", result, average, number)

	return average
}

func StorePrice(resp http.ResponseWriter, request *http.Request) {
	number, err := strconv.Atoi(request.FormValue("price"))
	if err != nil {
		log.Printf("%s\n", err)
	}
	numbers = append(numbers, number)

	_, err = io.WriteString(resp, fmt.Sprintf("%s", numbers))
	if err != nil {
		log.Printf("%s\n", err)
	}
	log.Printf("%d\n", numbers)
}
