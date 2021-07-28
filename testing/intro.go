package main

import "fmt"

func main() {
	f := foo(2, 4)
	g, err := bar(10.22)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(f)
	fmt.Println(g)
}

func foo(x, y int) int {
	return x + y
}

func bar(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("an error has occurred: %v", f)
	}
	return f, nil
}
