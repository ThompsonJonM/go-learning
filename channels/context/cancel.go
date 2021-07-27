package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("Error check:", ctx.Err())
	fmt.Println("Number of go routines:", runtime.NumGoroutine())

	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				n++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("Working...", n)
			}
		}
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("Error check 2:", ctx.Err())
	fmt.Println("Number of go routines:", runtime.NumGoroutine())

	fmt.Println("About to cancel...")
	cancel()
}
