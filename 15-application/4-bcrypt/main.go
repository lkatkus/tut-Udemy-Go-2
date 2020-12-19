package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	ps := "veryGoodPassword"

	bs, err := bcrypt.GenerateFromPassword([]byte(ps), bcrypt.MinCost)

	if err != nil {
		fmt.Println(err)
	}

	loginPassword := "veryGoodPasswords"

	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPassword))

	if err != nil {
		fmt.Println("Check your password")
	} else {
		fmt.Println("You're logged in")
	}
}
