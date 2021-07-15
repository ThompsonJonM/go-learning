package main

import "fmt"

func main() {
	// Composite literal: m := map[keyType]valueType{}
	m := map[string]int{
		"James": 32,
		"Alec":  45,
	}
	fmt.Println(m)
	fmt.Println(m["Alec"])

	v, ok := m["Barnabas"] // Comma OK idiom
	fmt.Println(v, ok)

	if v, ok := m["Barnabas"]; ok {
		fmt.Println("This is the if print", v) // Will not print as OK is false
	}
}
