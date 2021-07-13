package main

import "fmt"

func main() {
	x := 1986
	for {
		if x > 2021 {
			break
		}
		fmt.Println(x)
		x++
	}
}
