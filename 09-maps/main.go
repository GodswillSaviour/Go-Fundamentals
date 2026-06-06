// Project: Country Capital Finder
// Topic: Maps
// Description: Demonstrates the use of maps in Go

package main

import "fmt"

func main() {

	// ======= Declaring Map =======
	capitals := map[string]string{
		"Nigeria": "Abuja",
		"USA":     "Washington DC",
		"UK":      "London",
		"France":  "Paris",
		"Germany": "Berlin",
	}

	// ======= Taking User Input =======
	var country string
	fmt.Println("======= Country Capital Finder =======")
	fmt.Printf("Enter a country: ")
	fmt.Scanln(&country)
	fmt.Println("======================================")

	// ======= Searching Map =======
	capital, exists := capitals[country]
	if exists {
		fmt.Printf("Capital of %s is %s\n", country, capital)
	} else {
		fmt.Printf("%s not found in the list\n", country)
	}

	// ======= Printing All Countries =======
	fmt.Println("======= All Countries =======")
	for key, value := range capitals {
		fmt.Printf("%s: %s\n", key, value)
	}
	fmt.Println("=============================")
}