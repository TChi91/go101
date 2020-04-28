package main

import (
	"fmt"
	"io"
)

//MyString structure
type MyString struct {
	str       string
	readIndex int
}

func (data *MyString) Read(b []byte) (n int, err error) {
	//convert 'str' string to a slice of bytes
	strByte := []byte(data.str)

	//if 'readIndex' is GTE source length, return 'EOF' error
	if data.readIndex >= len(strByte) {
		return 0, io.EOF
	}

	//get next readable limit (exclusive)
	nextReadLimit := data.readIndex + len(b)

	//if 'nextReadLimit' is GTE source length
	//set it to source length and err to EOF
	if nextReadLimit >= len(strByte) {
		nextReadLimit = len(strByte)
		err = io.EOF
	}

	//get next bytes to copy and set 'n' to its length
	nextBytes := strByte[data.readIndex:nextReadLimit]
	n = len(nextBytes)

	//copy all bytes of nextBytes to B slice
	copy(b, nextBytes)

	//increment readIndex to nextReadLimit
	data.readIndex = nextReadLimit
	return
}

func main() {
	// create data source
	src := MyString{
		str: "Hello Amazing World!",
	}

	// create a packet
	p := make([]byte, 3) // slice of length `3`

	// read `src` until an error is returned
	for {

		// read `p` bytes from `src`
		n, err := src.Read(p)
		fmt.Printf("%d bytes read, data: %s\n", n, p[:])

		// handle error
		if err == io.EOF {
			fmt.Println("--end-of-file--")
			break
		} else if err != nil {
			fmt.Println("Oops! Some error occured!", err)
			break
		}
	}
}
