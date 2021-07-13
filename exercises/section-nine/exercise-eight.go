package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("This should not print.")
	case true:
		fmt.Println("This should print.")
	}
}
