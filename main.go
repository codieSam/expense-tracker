package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type Expense struct {
	ID          int       `json:"id"`
	Description string    `json:"desc"`
	Amount      float64   `json:"amount"`
	Date        time.Time `json:"time"`
}

var expenses []Expense

func loadExpenses() ([]Expense, error) {
	expenseData, err := os.ReadFile("expenses.json")
	if err != nil {
		fmt.Println("Error while reading the file", err)
		return nil, err
	}
	err = json.Unmarshal(expenseData, &expenses)
	if err != nil {
		fmt.Println("Error while decoding the json data.")
		return nil, err
	}
	return expenses, nil
}

func saveExpenses(expenses []Expense) error {
	jsonExpense, err := json.MarshalIndent(expenses, "", " ")
	if err != nil {
		fmt.Println("Error while marshelling the data")
		return err
	}
	err = os.WriteFile("expenses.json", jsonExpense, 0644)
	if err != nil {
		fmt.Println("Error while writing the file.")
		return err
	}
	return nil
}

func main() {
	fmt.Println("Welcome to expense-tracker")
	expenses, err := loadExpenses()
	if err != nil {
		fmt.Println("Error while loading the file")
		return
	}
	fmt.Println(expenses)

	var desc string
	var amount string
	if len(os.Args) == 6 && os.Args[1] == "add" && os.Args[2] == "--description" && os.Args[4] == "--amount" {
		desc = os.Args[3]
		amount = os.Args[5]
		fAmount, err := strconv.ParseFloat(amount, 64)
		if err != nil {
			fmt.Println("Error while parsing the amount to float64.")
			return
		}
		description := strings.TrimSpace(desc)
		if description == "" {
			fmt.Println("Please provide the full description.")
			return
		} else if fAmount <= 0 {
			fmt.Println("Please provide the positive and non-zero amount.")
			return
		} else if math.IsNaN(fAmount) == true {
			fmt.Println("Please provide the non NaN value of amount.")
			return
		} else if math.IsInf(fAmount, 0) == true {
			fmt.Println("Please provide the non infinite value of amount.")
			return
		}
		expense := Expense{
			ID:          1,
			Description: description,
			Amount:      fAmount,
			Date:        time.Now(),
		}
		// fmt.Println(expense)
		expenses = append(expenses, expense)
		saveExpenses(expenses)
	} else {
		fmt.Println(`Usage: go run . add --description "Grocery" --amount 20.32`)
		return
	}

}
