package PointerFunctions

import (
	"fmt"
)

// Person struct to create name and age of person
type person struct {
	name string
	age  int
}

// CreatePerson method creates a person struct
func CreatePerson(name string, age int) person {
	s := person{name: name, age: age}
	return s
}

// SetNameUsingPointer method sets value using pointer
func SetNameUsingPointer(personPointer *person) {
	personPointer.name = "Gopal"
}

//SetNameUsingValue method sets value of name by passing value.
func SetNameUsingValue(personPointer person) {
	personPointer.name = "Gopal2"
}

func (p *person) displayPersonInfo(messageText string) {
	fmt.Println(messageText, p.age, p.name)
}
