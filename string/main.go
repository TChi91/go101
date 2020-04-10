package main

import "fmt"

func main() {
	str := "Hello Fethi !"
	t := []rune(str)
	for i := 0; i < len(t); i++ {
		fmt.Printf("%c ", t[i])

	}
}

/*
%c ==> character
%v ==> value
%T ==> Type
%x ==> Hex Value
*/
