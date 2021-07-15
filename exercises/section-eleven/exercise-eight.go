package main

import "fmt"

func main() {
	m := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Ice cream", "Sunsets"},
	}

	fmt.Println(m)

	m["trevelyan_alec"] = []string{"Janus", "England", "Natalya"}

	fmt.Println(m)

	delete(m, "no_dr")

	for k, v := range m {
		fmt.Printf("Record for %v\n", k)
		for i, val := range v {
			fmt.Printf("\tThe index is %v and value is %v\n", i, val)
		}
	}
}
