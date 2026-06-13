// Project: Shape Calculator
// Topic: Interfaces
// Description: Demonstrates the use of interfaces in Go

package main

import "fmt"

// ======= Structs =======
type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

// ======= Rectangle Methods =======
func (r Rectangle) area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// ======= Circle Methods =======
func (c Circle) area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

// ======= Interface =======
type Shape interface {
	area() float64
	perimeter() float64
}

// ======= Interface Function =======
func printShapeInfo(s Shape) {
	fmt.Printf("  Area:      %.2f\n", s.area())
	fmt.Printf("  Perimeter: %.2f\n", s.perimeter())
}

func main() {

	// ======= Creating Shapes =======
	rect := Rectangle{Width: 5, Height: 3}
	circle := Circle{Radius: 2}

	// ======= Printing Shape Info =======
	fmt.Println("======= Shape Calculator =======")
	fmt.Println("Rectangle:")
	fmt.Printf("  Width:  %.2f\n", rect.Width)
	fmt.Printf("  Height: %.2f\n", rect.Height)
	printShapeInfo(rect)
	fmt.Println("--------------------------------")
	fmt.Println("Circle:")
	fmt.Printf("  Radius: %.2f\n", circle.Radius)
	printShapeInfo(circle)
	fmt.Println("================================")
}