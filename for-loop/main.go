//The for loop is the only loop statement in Go.

/*
for initialization; condition; post {
	// ... statments
}
*/

/*
// a traditional "while" loop
for condition {
	// ...
}
*/

/*
// a traditional infinite loop
for {
	// ...
}
*/
package main

import (
	"fmt"
)

var word string

//Standard ForLoop
func Standard() {
	word = "Hello from Standard!"
	for i := 0; i < 5; i++ {
		fmt.Println(word)
	}
}

// ForLoop "for item in iterable object"
func ForLoop() {
	var FullName = []string{"DECHAICHA ", "Fethi "}
	for _, name := range FullName {
		fmt.Printf(name)
	}
	// make new line
	fmt.Print("\n")
}

// ForLoopIndex "for item in iterable object"
func ForLoopIndex() {
	var FullName = []string{"DECHAICHA", "Fethi"}
	for index, name := range FullName {
		fmt.Println("index =", index, "and the name =", name)
	}
}

func main() {
	Standard()
	ForLoop()
	ForLoopIndex()
}
