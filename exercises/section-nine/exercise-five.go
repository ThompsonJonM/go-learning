package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		fmt.Printf("When %v is divded by 4, the remainder is %v\n", i, i%4)
	}
}
