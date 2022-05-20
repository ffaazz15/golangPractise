package main

import (
	"fmt"
)

func ex5() {
	var sum int
	var num int
	fmt.Println("enter the no")
	fmt.Scanln(&num)
	fmt.Println("your no is ", num)

	for num > 0 || sum > 9 {
		if num == 0 {
			num = sum
			sum = 0
		}
		sum += num % 10
		num /= 10
	}
	fmt.Println("your answer is is", sum)

}
