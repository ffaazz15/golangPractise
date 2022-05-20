package main

import (
	"fmt"
	"strings"
)

func ex2() {

	var str string
	var substr string
	var n int
	fmt.Printf("enter how many strings u want to :-")
	fmt.Scanf("%d", &n)
	for i:=0; i<n;i++{
	fmt.Println("enter the String:-")
	fmt.Scanf("%s", &str)
	fmt.Println("enter the SubString:-")
	fmt.Scanf("%s", &substr)

	fmt.Println("the string & substring are:", str, substr)
	if strings.Contains(str, substr) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}

	// str := "abcdefgh"
	// substr := "cd"

}
