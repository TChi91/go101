package main

import "fmt"

/* The make built-in function allocates and initializes an object of type slice, map, or chan (only).
Like new, the first argument is a type, not a value. Unlike new,
make's return type is the same as the type of its argument,
not a pointer to it. The specification of the result depends on the type.
*/

//Vertex type
// You have to initialize the map using the make function
// (or a map literal) before you can add any elements:
// This is the Map literals
type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

// or it can be like this:
/*
var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}
*/

func main() {
	// a slice of length 0 and capacity 10 is created
	names := make([]string, 0, 10)
	fmt.Println(names)

	fmt.Println("*********************")

	fmt.Println(m)

	fmt.Println("*********************")

	// Check if the elem exists in the map with a two-value assignment:
	elem, ok := m["Google"]
	if ok {
		fmt.Println(elem)
	} else {
		fmt.Println("The Vertex does not exists")
	}

}
