// Project: Bank Account Manager
// Topic: Pointers
// Description: Demonstrates the use of pointers in Go

package main

import "fmt"

// ======= Struct =======
type BankAccount struct {
	Name     string
	Balance  float64
	DarkMode bool
}

// ======= Functions =======
func deposit(account *BankAccount, amount float64) {
	account.Balance += amount
}

func withdraw(account *BankAccount, amount float64) {
	account.Balance -= amount
}

func updateName(account *BankAccount, newName string) {
	account.Name = newName
}

func toggleDarkMode(account *BankAccount) {
	account.DarkMode = !account.DarkMode
}

func main() {

	// ======= Creating Account =======
	account := BankAccount{Name: "Godswill", Balance: 5000.0, DarkMode: false}

	// ======= Displaying Opening Details =======
	fmt.Println("======= Bank Account Manager =======")
	fmt.Printf("Account Holder:  %s\n", account.Name)
	fmt.Printf("Opening Balance: %.2f\n", account.Balance)
	fmt.Println("====================================")

	// ======= Performing Operations =======
	fmt.Println("Depositing 2000.00...")
	deposit(&account, 2000.0)

	fmt.Println("Withdrawing 1000.00...")
	withdraw(&account, 1000.0)

	fmt.Println("Updating account holder name...")
	updateName(&account, "GodswillCodes")

	fmt.Println("Toggling dark mode...")
	toggleDarkMode(&account)
	fmt.Println("====================================")

	// ======= Displaying Account Summary =======
	fmt.Println("======= Account Summary =======")
	fmt.Printf("Account Holder: %s\n", account.Name)
	fmt.Printf("Balance:        %.2f\n", account.Balance)
	fmt.Printf("Dark Mode:      %t\n", account.DarkMode)
	fmt.Println("===============================")
}