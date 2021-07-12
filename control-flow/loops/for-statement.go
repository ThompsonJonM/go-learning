package main

import "fmt"

func main() {
	x := 1

	for x < 10 { // for statement with a single condition
		fmt.Println(x)
		x++
	}

	// for {
	// 	if x > 9 {
	// 		break
	// 	}

	// 	fmt.Println(x)
	// 	x++
	// }

	fmt.Println("Done.")
}
