package main

import "fmt"

// TODO: Create a user defined type "Price" which is an int64

type Product struct {
	Price Price
}

// TODO: Implement a *method* "Add" with Price as the receiver, which calculates the Sum of two Prices
/* Example

p1 := Price(123)
p2 := Price(100)
p3 := p1.Add(p2) // has the value 223

*/

// TODO: Implement a *method* "ToDecimal" with Price as the receiver, which converts the cents into euros
// Example: 199 => 1.99

// TODO: Implement a *method* "ToString" with Price as the receiver, which returns a string format of the price.
// Example: 199 => 1.99 EUR

// TODO: Write a function ToString for Price type that prints out the price in "12.39 EUR" format

func main() {
	mentosMint := Product{
		Price: 199,
	}
	mentosFruit := Product{
		Price: 149,
	}

	if mentosMint.Price.ToString() != "1.99 EUR" {
		panic("Wrong")
	}

	if mentosFruit.Price.ToString() != "1.49 EUR" {
		panic("Wrong")
	}

	if mentosMint.Price.Add(mentosFruit.Price).ToString() != "3.48 EUR" {
		panic("Wrong")
	}

	fmt.Println("Correct ✅")
}
