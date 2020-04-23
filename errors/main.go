package main

import (
	"errors"
	"fmt"
	"log"
)

type myerr struct{}

func (e *myerr) Error() string {
	return "my err"
}

//GreetFethi just a simple func
func GreetFethi(name string) (string, error) {
	if name != "Fethi" {
		return "", errors.New("i can greet just Fethi")
	}
	{
		greeting := fmt.Sprint("Hello from inside MyFunc")
		return greeting, nil
	}

}

func main() {
	errr := myerr{}
	fmt.Printf("%T\n", errr)

	variable, err := GreetFethi("Fethi")
	if err != nil {
		log.Fatal("Error: ", err)
	} else {
		fmt.Println(variable)
	}

	// You can also use Errorf function from fmt package to create interpolated error messages.
	e := errors.New("Error")
	e2 := fmt.Errorf("error %v with interpolation", e)
	fmt.Printf("%T\n", e)
	fmt.Printf("%T", e2)
}
