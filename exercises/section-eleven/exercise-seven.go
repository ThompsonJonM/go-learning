package main

import "fmt"

func main() {
	x := []string{"James", "Bond"}
	y := []string{"Miss", "Moneypenny"}
	z := [][]string{x, y}

	fmt.Println(z)

	for i, v := range z {
		fmt.Printf("Record %v\n", i)
		for _, v2 := range v {
			fmt.Printf("\tValues are %v\n", v2)
		}
	}
}
