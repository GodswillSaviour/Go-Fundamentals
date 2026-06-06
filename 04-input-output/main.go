// Project: User Info Form
// Topic: Input/Output
// Description: Demonstrates taking user input and displaying it

package main

import "fmt"

func main() {

	// ======= Declaring Variables =======
	var name string
	var age int
	var country string

	// ======= Taking User Input =======
	fmt.Println("======= User Info Form =======")
	fmt.Printf("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Printf("Enter your age: ")
	fmt.Scanln(&age)
	fmt.Printf("Enter your country: ")
	fmt.Scanln(&country)
	fmt.Println("==============================")

	// ======= Displaying User Info =======
	fmt.Println("======= Your Info =======")
	fmt.Printf("Name:    %s\n", name)
	fmt.Printf("Age:     %d\n", age)
	fmt.Printf("Country: %s\n", country)
	fmt.Println("=========================")
}