package main

import (
	"fmt"
	"strconv"

	"indramahesa.dev/gobank/fileops"
)

func main() {
	var accountBalance float64 = fileops.ReadBalanceToFile("data.json")
	var choice string

	fileops.WriteBalanceToFile("data.json", accountBalance)

	initialMessage()

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
			fileops.WriteBalanceToFile("data.json", accountBalance)
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
			fileops.WriteBalanceToFile("data.json", accountBalance)
			fmt.Println("Balance updated! New amount: " + strconv.Itoa(int(accountBalance)))
		case "exit":
			return
		case "help":
			initialMessage()
		default:
			fmt.Println("Menu not found!")
		}

		fmt.Println()
	}

}
