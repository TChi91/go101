package main

import "fmt"

func main() {
	var xx int = 2020
	// The & is called "ampersand operator"
	// The * is clled "dereferencing operator"
	fmt.Println("&xx ", *&xx)

	pp := &xx
	po := xx
	fmt.Println("pp ", *pp)
	*pp = 10
	fmt.Println(po, *pp)

	yy := new(int)
	fmt.Println(*yy)
	fmt.Printf("the value of yy is %v and the type is %T", *yy, yy)

}
