// Project: Grade Checker
// Topic: Conditionals
// Description: Demonstrates the use of if, else if and else in Go

package main

import "fmt"

func main() {

	// ======= Declaring Variable =======
	var score int
	var grade string
	var remark string

	// ======= Taking User Input =======
	fmt.Println("======= Grade Checker =======")
	fmt.Printf("Enter your score: ")
	fmt.Scanln(&score)
	fmt.Println("=============================")

	// ======= Checking Grade =======
	if score >= 90 {
		grade = "A"
		remark = "Excellent!"
	} else if score >= 80 {
		grade = "B"
		remark = "Very Good!"
	} else if score >= 70 {
		grade = "C"
		remark = "Good!"
	} else if score >= 60 {
		grade = "D"
		remark = "Pass!"
	} else {
		grade = "F"
		remark = "Failed!"
	}

	// ======= Displaying Result =======
	fmt.Println("======= Your Result =======")
	fmt.Printf("Score:  %d\n", score)
	fmt.Printf("Grade:  %s\n", grade)
	fmt.Printf("Remark: %s\n", remark)
	fmt.Println("===========================")
}