package main

import "fmt"

/*
fmt.Print(args…) (n int, e error)
The Print is a variadic function,
means it can accept multiple arguments.
It concatenates these arguments with space in between
if both of the adjacent argument is not a string
and writes to the standard output without a trailing newline.
*/
func main() {
	fmt.Print("Hello ", "World", "!")
	fmt.Println("The same as Print but add newline at the end")
	str := fmt.Sprint(`The same as Print and Println,
	but it returns the string and do not write it to stdout`)
	// %v is called a verb or placeholder
	// %#v is called: syntax representation
	// %+v display field names in a struct or pointer to a struct
	// # and + called "flags"
	// %T (Type format)
	// %t (boolean format)
	// %d (base10 format)
	// %b (base2/binary number format)
	// %x (hexadecimal format) and %o (octal format)
	// %s (string format) and %q (escaped string format)
	//  ...etc
	// For more about fmt https://golang.org/pkg/fmt/
	fmt.Printf("Here is str value: %v \n", str)
	fmt.Println("There in no Printfln, but yes there is Sprintf")

}
