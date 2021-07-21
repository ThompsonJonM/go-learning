package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var c int
var p int = 100

func main() {
	wg.Add(p)

	incrementer(p)

	wg.Wait()

	fmt.Println(runtime.NumGoroutine())
	fmt.Println("Counter:", c)
}

func incrementer(n int) {
	for i := 0; i < n; i++ {
		go func() {
			v := c
			runtime.Gosched()
			v++
			c = v
			wg.Done()
		}()
	}
}
