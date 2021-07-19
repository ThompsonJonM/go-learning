package main

import (
	"fmt"
	"sort"
)

func main() {
	x := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	y := []string{"James", "Q", "M", "Moneypenny", "Dr. No"}

	fmt.Println(x, y)

	sort.Ints(x)
	fmt.Println(x) // [3 4 7 12 16 18 42 56 99]
	fmt.Println(sort.IntsAreSorted(x)) // true

	sort.Strings(y)
	fmt.Println(y) // [Dr. No James M Moneypenny Q]
}
