package main

import "fmt"

const (
	a = iota
	b
	c
)

func main() {
	fmt.Println(a, b, c) // 0 1 2
}
