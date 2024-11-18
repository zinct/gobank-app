package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Bank struct {
	Balance float64
}

func writeBalanceToFile(accountBalance float64) {
	jsonString := `{"balance": ` + strconv.Itoa(int(accountBalance)) + `}`
	err := os.WriteFile("data.json", []byte(jsonString), 0644)

	if err != nil {
		fmt.Println("Failed to write balance: " + err.Error())
	}
}

func readBalanceToFile() float64 {
	data, err := os.ReadFile("data.json")

	if err != nil {
		return 0
	}

	var bank Bank

	err = json.Unmarshal(data, &bank)

	if err != nil {
		return 0
	}

	return bank.Balance
}

func main() {
	var accountBalance float64 = readBalanceToFile()
	var choice string

	writeBalanceToFile(accountBalance)

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println()

	for choice != "4" {
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case "1":
			fmt.Println("Your current balance is Rp." + strconv.Itoa(int(accountBalance)))
		case "2":
			var depositAmount int
			fmt.Print("Your deposit amount: ")
			fmt.Scan(&depositAmount)

			// Check if amount is greater than 0
			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must greater than 0")
				break
			}

			accountBalance += float64(depositAmount)
			writeBalanceToFile(accountBalance)
			fmt.Println("Balance updated! New amount: " + strconv.Itoa(int(accountBalance)))
		case "3":
			var withdrawAmount int
			fmt.Print("Your withdraw amount: ")
			fmt.Scan(&withdrawAmount)

			// Check if amount is greater than 0 and Insuffient to the balance
			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must greater than 0")
				break
			}

			if withdrawAmount > int(accountBalance) {
				fmt.Println("Invalid amount. Insuffient balance")
				break
			}

			accountBalance -= float64(withdrawAmount)
			writeBalanceToFile(accountBalance)
			fmt.Println("Balance updated! New amount: " + strconv.Itoa(int(accountBalance)))
		case "exit":
			return
		case "help":
			fmt.Println("Welcome to Go Bank!")
			fmt.Println("What do you want to do?")
			fmt.Println("1. Check balance")
			fmt.Println("2. Deposit Money")
			fmt.Println("3. Withdraw Money")
		default:
			fmt.Println("Menu not found!")
		}

		fmt.Println()
	}

}
