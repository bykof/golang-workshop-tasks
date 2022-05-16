package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

type Customer struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"first"`
	LastName  string    `json:"last"`
	Company   string    `json:"company"`
	CreatedAt time.Time `json:"created_at"`
	Country   string    `json:"country"`
}

type Customers []Customer

func (c Customer) String() string {
	return fmt.Sprintf("(%d) %s %s", c.ID, c.FirstName, c.LastName)
}

func (c Customers) Search(value string) Customers {
	var foundCustomers Customers
	for _, customer := range c {
		if strings.Contains(strings.ToLower(customer.FirstName), strings.ToLower(value)) {
			foundCustomers = append(foundCustomers, customer)
		}

		if strings.Contains(strings.ToLower(customer.LastName), strings.ToLower(value)) {
			foundCustomers = append(foundCustomers, customer)
		}

		if strings.Contains(strings.ToLower(customer.Company), strings.ToLower(value)) {
			foundCustomers = append(foundCustomers, customer)
		}

		if strings.Contains(strings.ToLower(customer.Email), strings.ToLower(value)) {
			foundCustomers = append(foundCustomers, customer)
		}
	}
	return foundCustomers
}

func (c Customers) FilterByCountry(value string) Customers {
	var filteredCustomers Customers
	for _, customer := range c {
		if customer.Country == value {
			filteredCustomers = append(filteredCustomers, customer)
		}
	}
	return filteredCustomers
}

func LoadCustomers(filePath string) Customers {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	// Now let's unmarshall the data into `payload`
	var customers Customers
	err = json.Unmarshal(content, &customers)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	return customers
}

/*
	Task:
		Create a Customer struct
		Create a Customers user-defined type which is type of []Customers
		Parse the "data.json" file with the encoding/json library
		Create a String() representation of a Customer
		Create a Search(value string) Customer method for Customers where you search for email, firstname, lastname and company
		Create a FilterByCountry(value string) Customers method, where you filter all customers by the given country
		Try to structure your application properly with functions so that your main becomes small! :)
*/
func main() {
	customers := LoadCustomers("./data.json")
	fmt.Printf("Found %d customers with a @hotmail.com email address\n", len(customers.Search("hotmail.com")))
	fmt.Printf("Found %d customers from the country \"Germany\"\n", len(customers.FilterByCountry("Germany")))
}
