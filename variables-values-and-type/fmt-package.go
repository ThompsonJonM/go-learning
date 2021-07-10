package main

import "fmt"

var y = 42

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%d\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%d\t%b\t%#x\n", y, y, y)

	s := fmt.Sprintf("%d\t%b", y, y)
	fmt.Println(s)
}
