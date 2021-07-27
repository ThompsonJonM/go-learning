package main

import (
	"fmt"
	"runtime"
)

func main() {
	// Channels allow us to pass values between goroutines
	c := make(chan int) // Channel which can take integers

	go func() {
		c <- 42                             // Place the value into the channel
		fmt.Println(runtime.NumGoroutine()) // 2
	}()

	fmt.Println(<-c) // Surface the value from the channel
}
