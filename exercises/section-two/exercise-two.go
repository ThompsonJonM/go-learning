package main

import "fmt"

var x int = 42
var y int = 50

func main() {
	e := x == y
	le := x <= y
	ge := x >= y
	ne := x != y
	lt := x < y
	gt := x > y

	fmt.Println(e)
	fmt.Println(le)
	fmt.Println(ge)
	fmt.Println(ne)
	fmt.Println(lt)
	fmt.Println(gt)
}
