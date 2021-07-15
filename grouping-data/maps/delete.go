package main

import "fmt"

func main() {
	m := map[string]int{
		"Jonathan": 34,
		"Nicole":   35,
	}

	fmt.Println(m)

	delete(m, "Jonathan")
	delete(m, "James") // No error thrown

	fmt.Println(m)

	if v, ok := m["Jonathan"]; ok {
		fmt.Println(v)
	}
}
