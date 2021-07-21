package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup // Establish wait group

func main() {
	wg.Add(2)

	go foo()
	go bar()

	fmt.Println(runtime.NumGoroutine())

	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("Foo Count:", i)
	}

	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("Bar Count:", i)
	}

	wg.Done()
}
