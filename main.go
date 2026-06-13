package main

import "fmt"

type Animal interface {
    speak() string
}

type Dog struct{}
type Cat struct{}

func (d Dog) speak() string {
    return "Woof!"
}

func (c Cat) speak() string {
    return "Meow!"
}

func makeSound(a Animal) {
    fmt.Println(a.speak())
}

func main() {
    dog := Dog{}
    cat := Cat{}

    makeSound(dog)
    makeSound(cat)
}