package main

import (
	"fmt"
	"time"
)

func goro_1() {
	fmt.Println("starts")

	go say_hi(1)
	go say_hi(2)
	go say_hi(3)

	fmt.Println("ends")
	time.Sleep(100 * time.Millisecond)

}
func say_hi(group int) {
	for i := 0; i < 3; i++ {
		fmt.Println(group, ":hello!")
	}
}
