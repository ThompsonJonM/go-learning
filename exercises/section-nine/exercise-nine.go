package main

import "fmt"

func main() {
	favSport := "football"

	switch favSport {
	case "rugby":
		fmt.Println("My favorite sport is rugby")
	case "football":
		fmt.Println("My favorite sport is football")
	case "soccer":
		fmt.Println("My favorite sport is rugby")
	}
}
