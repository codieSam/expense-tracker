# Go Expense Tracker CLI

A command-line Expense Tracker built with **Go (Golang)** to manage expenses, track spending by category, calculate monthly totals, and monitor a monthly budget.

This project was developed as a hands-on learning experience to practise Go fundamentals, JSON file persistence, input validation, error handling, and command-line application development.

## Features

- **Expense management:** Add, list, update, and delete expenses.
- **Expense summaries:** Calculate total expenditure and summaries by month or category.
- **Category tracking:** Organise expenses into categories such as Food, Transport, and Education.
- **Monthly budget management:** Set and save a monthly spending limit.
- **Budget monitoring:** Calculate monthly spending, remaining budget, and display warnings when the budget is reached or exceeded.
- **JSON persistence:** Store expense records and budget settings locally.
- **Input validation:** Handle invalid amounts, months, IDs, and other incorrect inputs.

## Tech Stack

- **Language:** Go (Golang)
- **Data storage:** JSON files
- **Version control:** Git and GitHub

## Prerequisites

Install [Go](https://go.dev/dl/) and verify the installation:


go version


## Getting Started

Clone the repository and navigate into the project directory:


git clone https://github.com/codieSam/expense-tracker.git
cd expense-tracker


Run the application directly:


go run .


You can also build a standalone executable:


go build -o expense-tracker.exe .


Run the executable on Windows:


.\expense-tracker.exe list


## Usage

Run the following commands from the project directory.

### 1. Add an expense


go run . add --description "Lunch" --amount 8 --category Food


### 2. List all expenses


go run . list


### 3. Update an expense

go run . update --id 2 --description "Weekly groceries" --amount 35.50


### 4. Delete an expense


go run . delete --id 2


### 5. View total expenses

go run . summary


### 6. View expenses for a month

Use a month number between 1 and 12:


go run . summary --month 10

### 7. View expenses by category


go run . summary --category Food


### 8. Set a monthly budget


go run . set-budget --amount 500


### 9. Check the monthly budget


go run . budget


The budget command displays the configured limit, spending for the current month, the remaining budget, and a warning when spending reaches or exceeds the limit.

## Data Storage

The application uses local JSON files to persist data:

- `expenses.json` — expense records, including descriptions, amounts, categories, and dates.
- `monthly_budget.json` — the configured monthly budget limit.

These files are created or updated by the application as data is saved.

**Note:** Keep personal or test data out of public repositories. Consider adding these runtime data files and generated executables to `.gitignore`.

## Go Concepts Practised

This project provided practical experience with:

- Structs and slices
- Functions and range loops
- Command-line arguments using `os.Args`
- Type conversion using `strconv`
- JSON encoding and decoding
- File operations using `os`
- Date and time handling with `time`
- Input validation and error handling
- Git version control and code formatting

## Future Improvements

Potential enhancements include:

- Automated tests using Go's `testing` package
- Better command parsing and a cleaner project structure
- More robust data validation
- CSV export and reporting
- Improved documentation and user experience

## Learning Outcome

Building this project helped me apply Go fundamentals to a practical application and understand how data persistence, command handling, validation, and business logic work together in a small software project.


**Built with Go as part of my ongoing software development and backend engineering learning journey.**