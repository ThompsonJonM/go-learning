package main

import "fmt"

func main() {
	x := 41
	if x > 42 {
		fmt.Println("The value is greater than 42")
	} else if x < 42 {
		fmt.Println("The value is less than 42")
	} else {
		fmt.Printf("The value is: %v", x)
	}
}
