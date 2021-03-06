package main

import "fmt"

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// send
	go send(eve, odd, quit)
	receive(eve, odd, quit)

	fmt.Println("Exiting")
}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
}

func receive(e, o, q <-chan int) {
	for {
		select { // Pulls value ready to be pulled from channel
		case v := <-e:
			fmt.Println("From the even channel:", v)
		case v := <-o:
			fmt.Println("From the odd channel:", v)
		case v := <-q:
			fmt.Println("From the quit channel:", v)
			return
		}
	}
}
