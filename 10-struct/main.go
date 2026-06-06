// Project: Student Records
// Topic: Structs
// Description: Demonstrates the use of structs in Go

package main

import "fmt"

// ======= Declaring Struct =======
type Student struct {
	Name  string
	Age   int
	Score float64
	Grade string
}

func main() {

	// ======= Creating Slice of Students =======
	students := []Student{
		{Name: "Godswill", Age: 24, Score: 98.50, Grade: "A"},
		{Name: "John", Age: 22, Score: 75.00, Grade: "C"},
		{Name: "Sarah", Age: 23, Score: 85.00, Grade: "B"},
	}

	// ======= Printing Student Records =======
	fmt.Println("======= Student Records =======")
	for i, student := range students {
		fmt.Printf("Student %d:\n", i+1)
		fmt.Printf("  Name:  %s\n", student.Name)
		fmt.Printf("  Age:   %d\n", student.Age)
		fmt.Printf("  Score: %.2f\n", student.Score)
		fmt.Printf("  Grade: %s\n", student.Grade)
		fmt.Println("------------------------------")
	}
	fmt.Println("===============================")

	// ======= Total Students =======
	fmt.Printf("Total Students: %d\n", len(students))
}