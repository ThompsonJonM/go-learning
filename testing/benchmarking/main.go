package main

import (
	"fmt"
	"github.com/ThompsonJonM/go-learning/testing/benchmarking/mystring"
	"strings"
)

const s = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla tristique magna hendrerit, vehicula ante sed, gravida lorem. Praesent vel purus gravida, laoreet orci non, condimentum felis. Morbi blandit ex leo, eu vehicula dolor dapibus vitae dapibus nam."

func main() {
	xs := strings.Split(s, " ")

	for _, v := range xs {
		fmt.Println(v)
	}

	fmt.Println(mystring.Cat(xs))
	fmt.Println(mystring.Join(xs))
}
