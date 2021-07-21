package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("Go Routines\t", runtime.NumGoroutine())

	counter := 0
	var wg sync.WaitGroup
	const gs = 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("Go Routines\t", runtime.NumGoroutine())
	}

	wg.Wait()

	fmt.Println("Go Routines\t", runtime.NumGoroutine())
	fmt.Println("Counter:", counter) // Intermittent due to race condition
	// Run with -race flag to check for race conditions using go run
}
