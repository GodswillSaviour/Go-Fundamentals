// Project: Data Types Explorer
// Topic: Data Types
// Description: Demonstrates the different data types in Go

package main

import "fmt"

func main() {

	// ======= Declaring Variables with different data types =======

	// String type
	var fullName string = "Godswill"

	// Integer types
	var age int = 24
	var bigNumber int64 = 9223372036854775807

	// Float types
	var pi32 float32 = 3.14
	var pi64 float64 = 3.141592653589793

	// Boolean type
	var active bool = true

	// Byte type
	var letter byte = 'G'

	// ======= Printing Variables =======

	fmt.Println("======= Data Types Explorer =======")
	fmt.Printf("String:  %s\n", fullName)
	fmt.Printf("Int:     %d\n", age)
	fmt.Printf("Int64:   %d\n", bigNumber)
	fmt.Printf("Float32: %.2f\n", pi32)
	fmt.Printf("Float64: %.15f\n", pi64)
	fmt.Printf("Bool:    %t\n", active)
	fmt.Printf("Byte:    %d\n", letter)
	fmt.Println("===================================")
}