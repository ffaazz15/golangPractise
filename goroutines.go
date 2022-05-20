package main

import (
	"fmt"
	"time"
)

func goroutines() {
	var msg = "hey hiii"
	go func() {
		fmt.Println(msg)

	}()
	{
		msg = "good buyeeee"
		time.Sleep(100 * time.Millisecond)
	}
}
