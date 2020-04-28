package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	var str string = "Hello Jacob !"

	//returns a pointer to strings.Reader struct
	//that implements the Read method defined by the io.Reader interface
	strReader := strings.NewReader(str)

	//If wanna read all the source data

	data, err := ioutil.ReadAll(strReader)
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Printf("%s\n", data)

}
