package main

import "fmt"

type person struct {
	first  string
	last   string
	flavor string
}

func main() {
	p1 := person{
		first:  "Jonathan",
		last:   "Thompson",
		flavor: "chocolate",
	}

	p2 := person{
		first:  "Nicole",
		last:   "Thompson",
		flavor: "vanilla",
	}

	fmt.Println(p1, p2)
	for _, v := range []person{p1, p2} {
		fmt.Println(v.flavor)
	}
}
