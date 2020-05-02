package main

import (
	"encoding/json"
	"fmt"
)

//Student structure
type Student struct {
	FirstName, LastName string
	Class               string
	Age                 uint8
	Grade               string
	username            string
}

//map values can be any type because all types implement the empty interface
var mp = map[string]interface{}{
	"FirstName": "DECHAICHA",
	"LastName":  "Fethi",
	"Class":     "2AS1",
	"age":       "18",
}
var name string = "TchiTchi"

func main() {
	tchi := Student{
		FirstName: "DECHAICHA",
		LastName:  "Fethi",
		Class:     "2AS1",
		Age:       18,
	}

	tchi2 := Student{
		FirstName: "DECHAICHA",
		LastName:  "Fethi",
		Class:     "2AS1",
		Age:       18,
		username:  "tchi",
	}
	slice := []Student{}
	slice = append(slice, tchi2, tchi)
	enc, _ := json.MarshalIndent(mp, "", " ")
	enc3, _ := json.MarshalIndent(slice, "", " ")
	fmt.Printf("%s\n", enc)
	fmt.Println("**********************************")
	fmt.Printf("%s", enc3)
}
