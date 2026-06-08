// Project: Simple Calculator
// Topic: Functions
// Description: Demonstrates the use of functions in Go

package main

import "fmt"

// ======= Calculator Functions =======
func add(x, y float64) float64 {
	return x + y
}

func subtract(x, y float64) float64 {
	return x - y
}

func multiply(x, y float64) float64 {
	return x * y
}

func divide(x, y float64) float64 {
	return x / y
}

func main() {

	// ======= Taking User Input =======
	var num1, num2 float64
	var operator string

	fmt.Println("======= Simple Calculator =======")
	fmt.Printf("Enter first number: ")
	fmt.Scanln(&num1)
	fmt.Printf("Enter operator (+, -, *, /): ")
	fmt.Scanln(&operator)
	fmt.Printf("Enter second number: ")
	fmt.Scanln(&num2)
	fmt.Println("=================================")

	// ======= Calling the right function =======
	var result float64
	switch operator {
	case "+":
		result = add(num1, num2)
	case "-":
		result = subtract(num1, num2)
	case "*":
		result = multiply(num1, num2)
	case "/":
		result = divide(num1, num2)
	default:
		fmt.Println("Invalid operator!")
		return
	}

	// ======= Displaying Result =======
	fmt.Printf("%.2f %s %.2f = %.2f\n", num1, operator, num2, result)
}