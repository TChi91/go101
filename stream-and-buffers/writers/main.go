package main

import (
	"fmt"
	"io"
)

// SampleStore - sample store type
type SampleStore struct {
	data []byte
}

// implement `io.Writer` interface
func (ss *SampleStore) Write(p []byte) (n int, err error) {

	// check if `10` bytes has been written
	if len(ss.data) == 10 {
		return 0, io.EOF // end of limit error
	}

	// get remaining capacity of the `ss.data`
	remainingCap := 10 - len(ss.data)

	// get length of data to write
	writeLength := len(p)
	if remainingCap <= writeLength {
		writeLength = remainingCap
		err = io.EOF
	}

	// append `writeLength` of data from `p` to `ss.data`
	ss.data = append(ss.data, p[:writeLength]...)

	// set number of bytes written and return
	n = writeLength
	return
}

func main() {

	ss := SampleStore{}

	// write 1: "Hello!"
	bytesWritten1, err1 := ss.Write([]byte("Hello!"))
	fmt.Printf("Bytes written %d, error: %v\n", bytesWritten1, err1)
	fmt.Printf("Value of ss.data: %s\n\n", ss.data)

	// write 2: " Amazing"
	bytesWritten2, err2 := ss.Write([]byte(" Amazing"))
	fmt.Printf("Bytes written %d, error: %v\n", bytesWritten2, err2)
	fmt.Printf("Value of ss.data: %s\n\n", ss.data)

	// write 3: " World!"
	bytesWritten3, err3 := ss.Write([]byte(" World!"))
	fmt.Printf("Bytes written %d, error: %v\n", bytesWritten3, err3)
	fmt.Printf("Value of ss.data: %s\n\n", ss.data)

	// Using the io.WriteString method
	bytesWritten4, err4 := io.WriteString(&ss, "ssss")
	fmt.Printf("Bytes written %d, error: %v\n", bytesWritten4, err4)
	fmt.Printf("Value of ss.data: %s\n\n", ss.data)

}
