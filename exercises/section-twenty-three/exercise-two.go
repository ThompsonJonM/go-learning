package main

import "fmt"

func main() {
	cs := make(chan int)
	cr := make(chan int)

	go func() {
		cs <- 42
		cr <- 42
	}()

	fmt.Println(<-cs)
	fmt.Println(<-cr)
}
