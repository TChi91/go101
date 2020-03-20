//Struct types
/*
A struct is a sequence of named elements, called fields,
each of which has a name and a type.
Field names may be specified explicitly (IdentifierList) or implicitly (EmbeddedField).
Within a struct, non-blank field names must be unique.
*/

package main

import (
	"fmt"
)

// Person struct type
type Person struct {
	First       string
	Last        string
	YearOfBirth int
}

//Age of a person
// Age here is a method
func (p Person) Age() int {
	var ActualYear int = 2020
	return ActualYear - p.YearOfBirth
}

func main() {
	fethi := Person{"Fethi", "TChi", 1991}

	fmt.Println("Hello Fethi, your Age is:", fethi.Age())
}

/*

Methods are functions
Remember: a method is just a function with a receiver argument.
Here's Abs written as a regular function with no change in functionality.
func Age(p Person) int {
	var ActualYear int = 2020
	return ActualYear - p.YearOfBirth
}
*/
