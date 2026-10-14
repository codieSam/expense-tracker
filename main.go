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

func nextFunctionId(expenses []Expense) int {
	maxID := 0

	for _, expense := range expenses {
		if expense.ID > maxID {
			maxID = expense.ID
		}

	}
	return maxID + 1
}

func findExpenseIndex(expenses []Expense, id int) int {
	for i, expense := range expenses {
		if expense.ID == id {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println("Welcome to expense-tracker")
	expenses, err := loadExpenses()
	if err != nil {
		fmt.Println("Error while loading the file")
		return
	}

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
		expense.ID = nextFunctionId(expenses)
		expenses = append(expenses, expense)
		err = saveExpenses(expenses)
		if err != nil {
			fmt.Println("Error while saving the expenses", err)
			return
		}
		fmt.Printf("Expense added successfully with (ID: %d) & (Description: %q)", expense.ID, expense.Description)
	} else if len(os.Args) == 2 {

		command := os.Args[1]
		if command == "list" {
			if len(expenses) == 0 {
				fmt.Println("No expenses found.")
				return
			}
			for i := range expenses {
				fmt.Printf("ID: %d | Description: %s | Amount: %.2f | Time: %s\n", expenses[i].ID, expenses[i].Description, expenses[i].Amount, expenses[i].Date.Format("2006-01-02"))
			}

		} else {
			fmt.Println("Command not found.")
		}
	} else if len(os.Args) == 8 {

		id, err := strconv.Atoi(os.Args[3])

		if err != nil {
			fmt.Println("Error while converting string ID into int", err)
			return
		}

		index := findExpenseIndex(expenses, id)
		// fmt.Println(index)
		if index == -1 {
			fmt.Println("Expense doesn't found.")
			return
		} else {
			amountToUpdate, err := strconv.ParseFloat(os.Args[7], 64)
			if err != nil {
				fmt.Println("Erreor while converting amountToUpdate to float", err)
				return
			}
			description := strings.TrimSpace(os.Args[5])
			if description == "" {
				fmt.Println("Please enter the valid description")
				return
			} else if amountToUpdate <= 0 {
				fmt.Println("Please enter the amount more than 0.")
				return
			} else if math.IsNaN(amountToUpdate) == true {
				fmt.Println("Please enter a valid amound")
				return
			} else if math.IsInf(amountToUpdate, 0) == true {
				fmt.Println("Please enter a valid amound")
				return
			}

			expenses[index].Description = os.Args[5]
			expenses[index].Amount = float64(amountToUpdate)
			err = saveExpenses(expenses)
			if err != nil {
				fmt.Println("Error while saving the updated expense.", err)
				return
			}
			fmt.Printf("Expense with id: %d is updated successfully", expenses[index].ID)
		}
	} else {
		fmt.Println(`Usage: go run . add --description "Grocery" --amount 20.32`)
		return
	}

}
