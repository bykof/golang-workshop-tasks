package main

import (
	"fmt"
	"time"
)

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

type Employee struct {
	ID               int64
	FirstName        string
	LastName         string
	Birthday         time.Time
	StartedAtCompany time.Time
}

func (e Employee) String() string {
	return fmt.Sprintf("%s %s", e.FirstName, e.LastName)
}

type Developer struct {
	Employee
	DeveloperType string
}

type Manager struct {
	Employee
	Department string
}

func (d Developer) DetailedString() string {
	return fmt.Sprintf("Developer(%s) %s", d.String(), d.DeveloperType)
}

func (m Manager) DetailedString() string {
	return fmt.Sprintf("Manager(%s) %s", m.String(), m.Department)
}

var _ DetailStringer = Developer{}
var _ DetailStringer = Manager{}

func main() {
	// Your code goes here
	developer := Developer{
		Employee: Employee{
			ID:               1,
			FirstName:        "Michael",
			LastName:         "Bykovski",
			Birthday:         time.Now(),
			StartedAtCompany: time.Now(),
		},
		DeveloperType: "Full Stack",
	}

	manager := Manager{
		Employee: Employee{
			ID:               1,
			FirstName:        "Peter",
			LastName:         "Lustig",
			Birthday:         time.Now(),
			StartedAtCompany: time.Now(),
		},
		Department: "IT",
	}
	fmt.Println(developer.String())
	fmt.Println(developer.DetailedString())

	fmt.Println(manager.String())
	fmt.Println(manager.DetailedString())
}
