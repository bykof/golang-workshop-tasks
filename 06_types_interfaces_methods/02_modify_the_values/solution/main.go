package main

import "fmt"

/*
	Task:
    	Create a method for the struct Person, which changes the LastName of it's current struct.
    	Create another method for the struct Person, to change the LastName of a copy of the struct and return the changed value
    	Create another method for the struct Person, to print out FullName() string

*/

type Person struct {
	FirstName string
	LastName  string
}

func (p *Person) ChangeLastName(value string) {
	p.LastName = value
}

func (p Person) ChangeLastNameCopy(value string) Person {
	p.LastName = value
	return p
}

func (p Person) FullName() {
	fmt.Printf("%s %s\n", p.FirstName, p.LastName)
}

func main() {
	// Your code goes here
	person := Person{
		FirstName: "Michael",
		LastName:  "Bykovski",
	}
	person.FullName()
	person = person.ChangeLastNameCopy("Test")
	person.FullName()
	person.ChangeLastName("AnotherTest")
	person.FullName()
}
