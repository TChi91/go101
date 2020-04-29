package main

import (
	"fmt"
	"os"
)

func main() {

	// open a `info.txt` file from CWD for read-write
	file, _ := os.OpenFile("tchi.txt", os.O_RDWR, 0744)

	// write a string
	if n, err := file.WriteString("Hello World!\n"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Success: %d bytes written\n", n)
	}

	// close file for any I/O operations
	defer file.Close()

	// write new string
	if n, err := file.WriteString("How are you?"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Success: %d bytes written\n", n)
	}

}
