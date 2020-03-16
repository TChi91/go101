package main

import "fmt"

var word string

func main() {
	word = "Hello Go!"
	for i := 0; i < 5; i++ {
		fmt.Println(word)
	}
}
