// Project: Multiplication Table
// Topic: Loops
// Description: Demonstrates the use of for loops in Go

package main

import "fmt"

func main() {

	// ======= Declaring Variable =======
	var number int

	// ======= Taking User Input =======
	fmt.Println("======= Multiplication Table =======")
	fmt.Printf("Enter a number: ")
	fmt.Scanln(&number)
	fmt.Println("====================================")

	// ======= Printing Multiplication Table =======
	fmt.Printf("======= Table of %d =======\n", number)
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", number, i, number*i)
	}
	fmt.Println("===========================")
}
