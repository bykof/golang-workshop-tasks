package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
}

func changeName(p *Person, firstName string, lastName string) {
	panic("Implement me")
}

// TODO: Implement without chaning the function signature
func sum(a *int64, b *int64) *int64 {
	panic("Implement me")
}

func main() {
	person := Person{
		FirstName: "Michael",
		LastName:  "Tester",
	}

	changeName(person, "Test", "Tester")

	// TODO: create a and b
	c := sum(a, b)
	fmt.Println(c)
}
