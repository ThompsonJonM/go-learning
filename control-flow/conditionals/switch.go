package main

import "fmt"

func main() {
	switch { // missing switch expression defaults to "true"
	case false:
		fmt.Println("This should not print.")
	case 2 == 4:
		fmt.Println("This also should not print.")
	case 3 == 3:
		fmt.Println("This should print.")
	case 4 == 4:
		fmt.Println("This should only print with fallthrough.")
	default:
		fmt.Println("This is default.")
	}

	switch "Found" {
	case "Found":
		fmt.Println("Found it.")
	case "Not found":
		fmt.Println("Not found.")
	}
}
