package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
}

type Car struct {
	Name         string
	Manufacturer string
}

var jsonBlob = []byte(`[
	{"Name": "WRX", "Manufacturer": "Subaru"},
	{"Name": "BRZ", "Manufacturer": "Subaru/Toyota"},
	{"Name": "Tacoma", "Manufacturer": "Toyota"}	
]`)

var cars []Car

func main() {
	group := ColorGroup{
		ID:     1,
		Name:   "Red",
		Colors: []string{"Maroon", "Magenta", "Burgundy"},
	}

	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Provide the ADDRESS of cars or else [] is returned
	// Can ONLY unmarshal to a pointer address
	err = json.Unmarshal(jsonBlob, &cars) 
	if err != nil {
		fmt.Println("Error:", err)
	}

	os.Stdout.Write(b)
	fmt.Printf("\n")
	fmt.Printf("%+v\n", cars)
}
