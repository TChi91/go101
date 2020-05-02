package main

import (
	"encoding/json"
	"fmt"
)

// Profile declares `Profile` structure
type Profile struct {
	Username  string
	followers int
	Grades    map[string]string
}

// Student declares `Student` structure
type Student struct {
	Username, lastName string
	Age                int
	Profile
	Languages []string
}

func main() {

	var john Student

	// define `john` struct
	john = Student{
		Username: "John",
		lastName: "Doe",
		Age:      21,
		Profile: Profile{
			Username:  "johndoe91",
			followers: 1975,
			//Grades:    map[string]string{"Math": "A", "Science": "A+"},
		},
		Languages: []string{"English", "French"},
	}

	// encode `john` as JSON
	johnJSON, err := json.MarshalIndent(john, "", "  ")

	// print JSON string
	fmt.Println(string(johnJSON), err)
}
