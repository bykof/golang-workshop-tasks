package main

import "fmt"

/*
TODO: Make the code work. Do not change the main function!
*/

type Namer struct {
	Name string
}

type Adder struct{}

func (_ Adder) Add(a float64, b float64) float64 {
	return a + b
}

type Producter struct{}

func (_ Producter) Product(a float64, b float64) float64 {
	return a * b
}

type Divider struct{}

func (_ Divider) Divide(a float64, b float64) float64 {
	return a / b
}

type Calculator struct{}

// Do not change the content of the main function!
func main() {
	calculator := Calculator{
		Namer: Namer{Name: "Test"},
	}

	if calculator.Add(1, 2) != 3 {
		panic("Wrong")
	}

	if calculator.Product(3, 3) != 9 {
		panic("Wrong")
	}

	if calculator.Divide(9, 3) != 3 {
		panic("Wrong")
	}

	if calculator.Name != "Test" {
		panic("Wrong")
	}
	fmt.Println("Correct ✅")
}
