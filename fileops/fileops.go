package fileops

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Bank struct {
	Balance float64
}

func WriteBalanceToFile(filename string, accountBalance float64) {
	jsonString := `{"balance": ` + strconv.Itoa(int(accountBalance)) + `}`
	err := os.WriteFile(filename, []byte(jsonString), 0644)

	if err != nil {
		fmt.Println("Failed to write balance: " + err.Error())
	}
}

func ReadBalanceToFile(filename string) float64 {
	data, err := os.ReadFile(filename)

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
