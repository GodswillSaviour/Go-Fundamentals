// Project: Shopping List
// Topic: Slices
// Description: Demonstrates the use of slices in Go

package main

import "fmt"

func main() {

	// ======= Declaring Slice =======
	items := []string{"Bread", "Eggs", "Milk", "Rice", "Butter"}

	// ======= Printing Shopping List =======
	fmt.Println("======= Shopping List =======")
	for i, item := range items {
		fmt.Printf("%d. %s\n", i+1, item)
	}
	fmt.Println("=============================")

	// ======= Removing Item =======
	removed := items[2]
	items = append(items[:2], items[3:]...)
	fmt.Printf("Item removed: %s\n", removed)

	// ======= Printing Updated List =======
	fmt.Println("======= Updated List =======")
	for i, item := range items {
		fmt.Printf("%d. %s\n", i+1, item)
	}
	fmt.Println("============================")

	// ======= Total Items =======
	fmt.Printf("Total Items: %d\n", len(items))
}