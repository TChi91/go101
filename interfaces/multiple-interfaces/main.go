package main

import (
	"fmt"
)

//Shape interface
type Shape interface {
	Area() float64
}

//Object interface
type Object interface {
	Volume() float64
}

//Skin interface
type Skin interface {
	color() string
}

//Cube structure
type Cube struct {
	side float64
}

//Area method
func (c Cube) Area() float64 {
	return c.side * c.side * 6
}

//Volume method
func (c Cube) Volume() float64 {
	return c.side * c.side * c.side
}
func main() {
	c := Cube{3}
	var s Shape = c
	var o Object = c

	fmt.Printf("The area of c is %v\n", s.Area())
	fmt.Printf("The volume of c is %v\n", o.Volume())

	fmt.Printf("The area of c is %v\n", s.(Cube).Volume())
	fmt.Printf("The volume of c is %v\n", o.(Cube).Area())

	//s.(Type) ==> Type assertion
	_, ok := s.(Skin)
	fmt.Println("check if dynamic type of c implements interface Skin ==>", ok)
}
