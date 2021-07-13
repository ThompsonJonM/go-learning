package main

import "fmt"

func main() {
	fmt.Printf("The value of true AND true is: %v\n", true && true)
	fmt.Printf("The value of true AND false is: %v\n", true && false)
	fmt.Printf("The value of true OR true is: %v\n", true || true)
	fmt.Printf("The value of true OR flase is: %v\n", true || true)
	fmt.Printf("The value of NOT true is: %v\n", !true)
	fmt.Printf("The value of NOT false is: %v", !false)
}
