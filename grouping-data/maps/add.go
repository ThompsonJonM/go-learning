package main

import "fmt"

func main() {
	m := map[string]int{
		"James":    32,
		"Jonathan": 34,
		"Timothy":  36,
	}
	fmt.Println(m)

	m["Sarah"] = 38

	fmt.Println(m)

	for k, v := range m {
		fmt.Printf("The key is: %v and value is: %v\n", k, v)
	}
}
