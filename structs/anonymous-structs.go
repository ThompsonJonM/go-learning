package main

import "fmt"

func main() {
	// Used when trying to maintain "code pollution", such as when a type
	// will only be used in one specific area
	// Rather than building a type, just use an anonymous struct
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	fmt.Println(p1)
}
