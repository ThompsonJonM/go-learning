package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		fmt.Println("Error occurred:", err)
		log.Println("Log error occurred:", err)
		log.Panicln("Panic error occurred:", err)
		// log.Fatalln("Fatal error occurred:", err)
		// panic(err)
	}
}
