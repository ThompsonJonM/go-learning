package main

import (
	"fmt"
	"strings"
)

const s = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla tristique magna hendrerit, vehicula ante sed, gravida lorem. Praesent vel purus gravida, laoreet orci non, condimentum felis. Morbi blandit ex leo, eu vehicula dolor dapibus vitae dapibus nam."

func main() {
	xs := strings.Split(s, " ")

	for _, v := range xs {
		fmt.Println(v)
	}
}
