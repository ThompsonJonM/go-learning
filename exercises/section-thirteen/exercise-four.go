package main

import "fmt"

func main() {
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "Jonathan",
		last:  "Thompson",
		age:   34,
	}

	fmt.Println(p1)
}
