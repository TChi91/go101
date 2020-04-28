package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	var str string = "Hello Amine !"

	//returns a pointer to strings.Reader struct
	//that implements the Read method defined by the io.Reader interface
	strReader := strings.NewReader(str)

	//If wanna read all the source data
	buf := make([]byte, 15)
	fmt.Println("======================")
	fmt.Println("io.ReadFull")
	//fmt.Println("======================")

	n, err := io.ReadFull(strReader, buf)
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Printf("%s\n", buf[:n])

	fmt.Println("======================")
	fmt.Println("io.LimitReader")
	fmt.Println("======================")

	var myString string = "Hello There!!"
	r := strings.NewReader(myString)
	buff := make([]byte, 3)

	limitedString := io.LimitReader(r, 10)

	for {
		n, err := limitedString.Read(buff)
		if err != nil {
			fmt.Println("Error", err)
			break
		}
		fmt.Printf("%d ==> %s\n", n, buff[:n])

	}
}
