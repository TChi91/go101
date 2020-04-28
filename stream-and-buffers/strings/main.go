package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

func main() {
	var str string = "Hello Jacob !"

	//returns a pointer to strings.Reader struct
	//that implements the Read method defined by the io.Reader interface
	strReader := strings.NewReader(str)

	p := make([]byte, 5)
	for {
		n, err := strReader.Read(p)
		if err == io.EOF {
			fmt.Println("End Of File!")
			break
		}
		if err != nil {
			fmt.Println("Error!", err)
		}

		fmt.Printf("%d, %s\n", n, p[:n])

		//If wanna read all the source data

		data, err := ioutil.ReadAll(strReader)
		fmt.Printf("%s\n", data)

	}
}
