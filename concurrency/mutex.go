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

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock() // "Check out" the variable
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock() // "Check in" the variable
			wg.Done()
		}()
		fmt.Println("Go Routines\t", runtime.NumGoroutine())
	}

	wg.Wait()

	fmt.Println("Go Routines\t", runtime.NumGoroutine())
	fmt.Println("Counter:", counter) // 100
}
