package main

import "fmt"

func main() {
	c := make(chan int)
	// send
	go foo(c)
	// receive
	bar(c) // Removed 'go' so bar acts as a blocker until foo assigns a value

	fmt.Println("Exiting")
}

func foo(c chan<- int) {
	c <- 42
}

func bar(c <-chan int) {
	fmt.Println(<-c)
}
