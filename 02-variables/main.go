// Project: Personal Profile Card
// Topic: Variables
// Description: Demonstrates the use of variables in Go

package main

import "fmt"

func main() {

	// ======= Declaring Variables =======

	// Using var keyword for explicit type declaration
	var name string = "Godswill"
	var age int = 24
	var active bool = true

	// Using short variable declaration :=
	country := "Nigeria"
	programmingLanguage := "Go"
	score := 98.5

	// ======= Printing Variables =======

	fmt.Println("======= My Profile =======")
	fmt.Printf("Name:                %s\n", name)
	fmt.Printf("Age:                 %d\n", age)
	fmt.Printf("Country:             %s\n", country)
	fmt.Printf("Programming Language:%s\n", programmingLanguage)
	fmt.Printf("Score:               %.2f\n", score)
	fmt.Printf("Active:              %t\n", active)
	fmt.Println("==========================")
}