package main

import "fmt"

// DoSomething just print the hello world string
func DoSomething() {
	fmt.Println("Hello World!")
}

func withParam(s string) {
	fmt.Printf("The string in the func param is: %v \n", s)
}

// add return the sum of two integer
func add(x, y int) int64 {
	return int64(x + y)
}

// func in Go can return multiple value
func twoReturn(x, y int64) (xx int64, yy int64) {
	xx = x * 2
	yy = y * 2
	return
}
func main() {
	DoSomething()

	withParam("Hi There!!")

	sum := add(5, 20)
	fmt.Println("The sum is", sum)

	val1, val2 := twoReturn(20, 55)
	fmt.Println(val1, val2)
}
