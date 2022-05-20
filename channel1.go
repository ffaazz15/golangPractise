package main

import (
	"fmt"
	"sync"
)

var waitg = sync.WaitGroup{}

func channel1() {
	ch := make(chan int)
	waitg.Add(2)
	go func() {
		i := <-ch
		fmt.Println(i)
		waitg.Done()
	}()
	go func() {
		ch <- 42
		waitg.Done()
	}()
	waitg.Wait()
}
