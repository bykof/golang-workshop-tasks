package main

import "fmt"

/*
	Task:
		Create three structs:
			1. Employee
			2. Developer
			3. Manager

		An Employee has the following fields:
			- ID                int64
			- FirstName         string
			- LastName          string
			- Birthday          time.Time
			- StartedAtCompany  time.Time

		A Developer embeds an employee struct and has the following fields:
			- DeveloperType     string (frontend, backend, full stack)

		A Manager embeds an employee struct and has the following fields:
			- Department        string (IT, HR)

	Output:
		Michael Bykovski
		Michael Bykovski Full Stack
		Peter Lustig
		Peter Lustig IT
*/

type DetailStringer interface {
	fmt.Stringer
	DetailedString() string
}

// Your code goes here

func main() {
	// Your code goes here
	var _ DetailStringer = &Developer{}
	var _ DetailStringer = &Manager{}

	fmt.Println(developer.String())
	fmt.Println(developer.DetailedString())

	fmt.Println(manager.String())
	fmt.Println(manager.DetailedString())
}
