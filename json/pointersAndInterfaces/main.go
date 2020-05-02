package main

import (
	"encoding/json"
	"fmt"
)

// ProfileI interface defines `Follow` method
type ProfileI interface {
	Follow()
}

// Profile declares `Profile` structure
type Profile struct {
	Username  string
	Followers int
}

// Follow method implementation
func (p *Profile) Follow() {
	p.Followers++
}

// Student declares `Student` structure
type Student struct {
	FirstName, lastName string
	Age                 int
	Primary             ProfileI
	Secondary           ProfileI
}

func main() {
	var name string
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter text: ")
	// name, _ := reader.ReadString('\n')
	fmt.Scanln(&name)

	// define `john` struct (pointer)
	john := &Student{
		FirstName: name,
		lastName:  "Doe",
		Age:       21,
		Primary: &Profile{
			Username:  "johndoe91",
			Followers: 1975,
		},
	}

	// follow `john`
	john.Primary.Follow()

	// encode `john` as JSON
	johnJSON, _ := json.MarshalIndent(john, "", "  ")

	// print JSON string
	fmt.Println(string(johnJSON))

}
