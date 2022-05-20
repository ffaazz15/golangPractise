package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func say(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Microsecond * 100)
	}
	wg.Done()
}
func goro_2() {
	wg.Add(1)
	say("hi")
	wg.Add(1)
	say("fazil")

	wg.Wait()
}
