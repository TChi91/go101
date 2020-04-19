package main

import (
	"fmt"
	"math"
)

// Shape interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Triangle structure
type Triangle struct {
	base   float64
	height float64
}

/*
Since Area and Perimeter methods are defined by the Shape interface,
the Triangle struct type *"implements the Shape interface"*
*/

// Area method to calculate the area
func (t Triangle) Area() float64 {
	return t.base * t.height / 2
}

// Perimeter method to calculate the perimeter
func (t Triangle) Perimeter() float64 {
	const c float64 = 10
	return t.base + t.height + c
}

// Circle type
type Circle struct {
	radius float64
}

// Area method to calculate the area
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// Perimeter method to calculate the perimeter
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// Empty interface == accepts all types
func explain(i interface{}) {
	fmt.Printf("This method accepts all types -%T-", i)
}

func main() {
	// s var declared as Shape type but we can assign a struct of type Triangle
	var s Shape
	//This statement print <nil> because Printf show the dynamic type that we didn't assign yet
	fmt.Printf("s value is %v\n", s)
	fmt.Printf("s type is %T\n", s)

	s = Triangle{2.5, 5.3}

	// here Printf show the dynamic type main.Triangle
	/*
		💡 We call it dynamic because we can assign s with a new struct of a different struct type
		which also implements the interface Shape.
	*/
	fmt.Printf("s value is %v\n", s)
	fmt.Printf("s type is %T\n", s)

	// we can assign s that is Triangle to another  struct type "Circle"
	// because "Circle " also implement the Shape interface

	s = Circle{3}
	fmt.Printf("s value is %v\n", s)
	fmt.Printf("s type is %T\n", s)

	explain(s)
}
