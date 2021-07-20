package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	p := `password1234`
	bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(p)
	fmt.Println(bs)

	loginPassword := `password`
	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPassword))
	if err != nil {
		fmt.Println("Invalid password.")
	}
}