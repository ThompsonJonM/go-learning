package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Create a new wait group
var wg sync.WaitGroup

func main() {
	// Concurrency is a design pattern
	// Concurrent code can potentially run in parallel
	// Concurrency DOES NOT guarantee code will run in parallel (single CPU)
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("Go Routines\t", runtime.NumGoroutine())

	wg.Add(1) // Add one thing to wait for
	go foo()
	bar()

	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("Go Routines\t", runtime.NumGoroutine())
	wg.Wait() // Wait for the thing
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("Foo:", i)
	}

	wg.Done() // Signal that the thing is done
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("Bar:", i)
	}
}
