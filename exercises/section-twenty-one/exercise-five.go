package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var c int64
var p int = 100

func main() {
	wg.Add(p)

	incrementer(p)

	wg.Wait()

	fmt.Println(runtime.NumGoroutine())
	fmt.Println(c)
}

func incrementer(n int) {
	for i := 0; i < n; i++ {
		atomic.AddInt64(&c, 1)
		wg.Done()
	}
}
