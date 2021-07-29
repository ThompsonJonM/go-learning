package word

import (
	"fmt"
	"strings"
)

func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	fmt.Println(xs)
	m := make(map[string]int)

	for _, v := range xs {
		m[v]++
	}

	return m
}

func Count(s string) int {
	return len(s)
}
