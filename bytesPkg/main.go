package main

import (
	"bytes"
	"fmt"
)

func main() {
	bytesBuffer := new(bytes.Buffer)

	bytesBuffer.Write([]byte("Hello there!"))

	fmt.Fprintf(bytesBuffer, " Sa7iit frero\n")

	fmt.Println(bytesBuffer.Len())
	//bytesBuffer.WriteTo(os.Stdout)
}
