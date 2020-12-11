package main

import (
	"fmt"
	"log"

	"idGen/idGenerator"
)

func main() {
	idGenerator, err := idGenerator.New(1)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 100; i++ {
		fmt.Println(idGenerator.GetNextId())
	}
}
