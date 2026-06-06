// Project: Student Scores
// Topic: Arrays
// Description: Demonstrates the use of arrays in Go

package main

import "fmt"

func main() {

	// ======= Declaring Array =======
	var scores [5]int = [5]int{85, 90, 78, 92, 88}

	// ======= Declaring Variables =======
	highest := scores[0]
	lowest := scores[0]
	sum := 0

	// ======= Looping Through Array =======
	fmt.Println("======= Student Scores =======")
	for i := 0; i < len(scores); i++ {
		fmt.Printf("Student %d: %d\n", i+1, scores[i])
		sum += scores[i]
		if scores[i] > highest {
			highest = scores[i]
		}
		if scores[i] < lowest {
			lowest = scores[i]
		}
	}
	fmt.Println("==============================")

	// ======= Calculating Average =======
	average := float64(sum) / float64(len(scores))

	// ======= Displaying Results =======
	fmt.Printf("Highest Score: %d\n", highest)
	fmt.Printf("Lowest Score:  %d\n", lowest)
	fmt.Printf("Average Score: %.2f\n", average)
}
