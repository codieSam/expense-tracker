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
	Category    string    `json:"category"`
}

type MonthlyBudgetLimit struct {
	MonthlyBudgetLimit float64 `json:"monthly_limit"`
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

func nextExpenseId(expenses []Expense) int {
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

// var monthlyLimit []MonthlyBudgetLimit

func saveMonthlyBudgetLimit(monthlyLimit MonthlyBudgetLimit) error {
	encodedMonthlyBudget, err := json.MarshalIndent(monthlyLimit, "", "")
	if err != nil {
		fmt.Println("Error while marshelling budget limit.", err)
		return err
	}
	err = os.WriteFile("monthly_budget.json", encodedMonthlyBudget, 0644)
	if err != nil {
		fmt.Println("Error while writing the budget limit.", err)
		return err
	}
	fmt.Printf("Monthly budget saved successfully to %s.", os.Args[3])
	return nil
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
	if len(os.Args) == 8 && os.Args[1] == "add" && os.Args[2] == "--description" && os.Args[4] == "--amount" && os.Args[6] == "--category" {
		desc = os.Args[3]
		amount = os.Args[5]
		category := strings.TrimSpace(os.Args[7])
		if category == "" {
			fmt.Println("Please enter the category")
			return
		}
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
			Category:    category,
			Date:        time.Now(),
		}
		// fmt.Println(expense)
		expense.ID = nextExpenseId(expenses)
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
				category := expenses[i].Category
				if category == "" {
					category = "UnCategroized"
				}
				fmt.Printf("ID: %d | Description: %s | Amount: %.2f | Category: %s | Time: %s\n", expenses[i].ID, expenses[i].Description, expenses[i].Amount, category, expenses[i].Date.Format("2006-01-02"))
			}

		} else if command == "summary" {
			totalExpense := 0.0
			for _, expense := range expenses {
				totalExpense += expense.Amount
			}
			fmt.Printf("The total expenses is : $%.2f\n", totalExpense)
		}
	} else if len(os.Args) == 8 && os.Args[6] == "--update" {

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
	} else if len(os.Args) == 4 {
		if os.Args[1] == "set-budget" && os.Args[2] == "--amount" {
			limitAmount, err := strconv.ParseFloat(os.Args[3], 64)
			if err != nil {
				fmt.Println("Error  while converting limitamount to float64", err)
				return
			}

			monthlyLimit := MonthlyBudgetLimit{
				MonthlyBudgetLimit: limitAmount,
			}
			// fmt.Println(monthlyBudgetLimit)
			err = saveMonthlyBudgetLimit(monthlyLimit)

		} else if os.Args[1] == "delete" && os.Args[2] == "--id" {
			intID, err := strconv.Atoi(os.Args[3])
			if err != nil {
				fmt.Println("Error while covrting to int on delete function", err)
				return
			}
			index := findExpenseIndex(expenses, intID)
			if index == -1 {
				fmt.Printf("Expenses with id %d was not found.", intID)
				return
			}
			expenses = append(expenses[:index], expenses[index+1:]...)
			err = saveExpenses(expenses)
			if err != nil {
				fmt.Println("Error while saving the data after deletion", err)
				return
			}
			fmt.Printf("Expense with id:  %d is deleted.", intID)
		} else if os.Args[1] == "summary" && os.Args[2] == "--month" {
			givenMonth := os.Args[3]
			month, err := strconv.Atoi(givenMonth)
			if err != nil {
				fmt.Println("Error while converting month to  int.", err)
				return
			}
			if month < 1 || month > 12 {
				fmt.Println("Please enter a month between 1 & 12.")
				return
			}
			totalExpenseOfMonth := 0.0

			for _, expense := range expenses {
				if time.Month(month) == expense.Date.Month() {
					totalExpenseOfMonth += expense.Amount
				}

			}
			fmt.Printf("Total Expense of %s is: $%.2f\n", givenMonth, totalExpenseOfMonth)

		} else if os.Args[1] == "summary" && os.Args[2] == "--category" {
			category := os.Args[3]
			// fmt.Println(category)
			totalExpenseOfCategory := 0.0
			for _, expense := range expenses {
				if category == expense.Category {
					totalExpenseOfCategory += expense.Amount
				}
			}
			fmt.Printf("The summary for category %s is : %.2f \n ", category, totalExpenseOfCategory)
		}
	} else {
		fmt.Println("Please use correct command")
		return
	}

}
