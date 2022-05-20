package main

import (
	"fmt"
	"strings"
)

func oval() {

	var s string
	fmt.Println("enter the string")
	fmt.Scanln(&s)
	fmt.Printf("your string is %v \n ", s)

	if strings.ContainsAny(s, "aAeEiIoOuU")  {
		fmt.Printf("string buetiful \n")

	} else {
		fmt.Printf("not buetifoool!!")

	}

}

//strings.Contains("e", q[i])) && strings.Contains("i", q[i]) && strings.Contains("o", q[i]) && strings.Contains("u", q[i]) {
