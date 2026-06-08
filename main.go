package main

import "fmt"

func calculate(a, b int) (int, int) {
    sum := a + b
    difference := a - b
    return sum, difference
}

func main() {
    sum, difference := calculate(10, 3)
    fmt.Println(sum)        // 13
    fmt.Println(difference) // 7
}