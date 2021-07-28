package main

import (
	"fmt"
)

type customErr struct {
	info string
}

func (e customErr) Error() string {
	return fmt.Sprintf("an error occurred: %v\n", e.info)
}

func main() {
	e := customErr{
		info: "This should run",
	}

	foo(e)
}

func foo(e error) {
	fmt.Println(e)
}
