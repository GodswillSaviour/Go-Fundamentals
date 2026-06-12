// Project: Safe Calculator
// Topic: Error Handling
// Description: Demonstrates error handling in Go

package main

import "fmt"

// ======= Calculate Function with Error Handling =======
func calculate(a, b float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("cannot divide by zero")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("invalid operator: %s", operator)
	}
}

func main() {

	// ======= Taking User Input =======
	var a, b float64
	var operator string

	fmt.Println("======= Safe Calculator =======")
	fmt.Printf("Enter first number: ")
	fmt.Scanln(&a)
	fmt.Printf("Enter operator (+, -, *, /): ")
	fmt.Scanln(&operator)
	fmt.Printf("Enter second number: ")
	fmt.Scanln(&b)
	fmt.Println("===============================")

	// ======= Calculating Result =======
	result, err := calculate(a, b, operator)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Result: %.2f\n", result)
	}
}