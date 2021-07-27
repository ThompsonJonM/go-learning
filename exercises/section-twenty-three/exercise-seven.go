package main

import (
	"fmt"
	"runtime"
)

func main() {
	c := make(chan int)

	for i := 0; i < 10; i++ {
		fmt.Println("Number of go routines:", runtime.NumGoroutine())
		go func() {
			for j := 0; j < 10; j++ {
				c <- j
			}
		}()
	}

	for k := 0; k < 100; k++ {
		fmt.Println(k, <-c)
	}
}
