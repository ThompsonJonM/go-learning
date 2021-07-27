package main

import (
	"fmt"
)

func main() {
	c := make(chan<- int) // Send only channel
	d := make(<-chan int) // Receive only channel

	go func() {
		c <- 42
		// d <- 42 // Cannot send to receive only
	}()

	// fmt.Println(<- c) // Cannot receive from send only
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", d)
}
