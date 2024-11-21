package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
)

func initialMessage() {
	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach us on: " + randomdata.PhoneNumber())
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println()
}
