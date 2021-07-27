package main

import (
	"fmt"
	"runtime"
)

func main() {
	c := make(chan int, 1) // BUFFERED channel which allows a single int onto the channel

	go func() {
		c <- 42
		// c <- 43 // Leads to a channel deadlock as the buffered channel only allows a single int
		fmt.Println(runtime.NumGoroutine())
	}()

	fmt.Println(<-c)
}
