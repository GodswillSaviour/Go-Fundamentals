package main

import "fmt"

type Movie struct {
	Name   string
	Rating int
	Year   int
}

func (m Movie) getCategory() string {
if m.Rating >= 8 {
return "Excellent"
}else if m.Rating >= 5 {
return "Good"
}else{
return "Poor"
}
}

func main() {
movies := []Movie{
{"Inception", 9, 2010},
{"Avatar", 7, 2009},
{"Interstellar", 10, 2014},
{"Titanic", 6, 1997},
{"Unknown", 4, 2020},
}


for _, movie := range movies {
fmt.Printf("Name: %s, Year: %d, Rating: %d, Category: %s\n", movie.Name, movie.Year, movie.Rating, movie.getCategory())
}
}