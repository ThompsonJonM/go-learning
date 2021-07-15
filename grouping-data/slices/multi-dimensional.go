package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "P38"}
	fmt.Println(jb)

	at := []string{"Alec", "Trevelyan", "M16"}
	fmt.Println(at)

	xp := [][]string{jb, at} // Create a multi-dimensional slice
	fmt.Println(xp)
}
