package main

import "fmt"

func main() {

	var country string
	countries := make(map[string]string)
	fmt.Println("=======Country Capital Finder=======")
	fmt.Print("Enter a country name: ")
	fmt.Scanln(&country)
	fmt.Println("==================================")

	fmt.Printf("The capital of %s is: Abuja", country)

	fmt.Println("====== All Countries ======")
	countries["Nigeria"] = "Abuja"
	countries["USA"] = "Washington D.C."
	countries["UK"] = "London"
	countries["France"] = "Paris"
	countries["Germany"] = "Berlin"

	for key, value := range countries {
		fmt.Printf("%s: %s\n", key, value)
	}

}
