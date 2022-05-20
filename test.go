package main

import (
	"fmt"
	"log"
	"strings"
)

func test() {
	fmt.Println("input text:")
	var w1, w2, w3 string
	n, err := fmt.Scan(&w1, &w2, &w3)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("number of items read: %d\n", n)
	fmt.Printf("read text: %s %s %s\n", w1, w2, w3)

	if strings.ContainsAny(w1, "aAeEiIoOuU") {
		fmt.Printf("string buetiful \n")

	} else {
		fmt.Printf("not buetifoool!!")

	}
}
