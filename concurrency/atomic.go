package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("Go Routines\t", runtime.NumGoroutine())

	var counter int64
	var wg sync.WaitGroup
	const gs = 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("Counter:\t", atomic.LoadInt64(&counter))
			runtime.Gosched()
			wg.Done()
		}()
		fmt.Println("Go Routines\t", runtime.NumGoroutine())
	}

	wg.Wait()

	fmt.Println("Go Routines\t", runtime.NumGoroutine())
	fmt.Println("Counter:", counter) // 100
}
