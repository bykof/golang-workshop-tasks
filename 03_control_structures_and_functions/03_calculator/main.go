package main

import "fmt"

func doOperation(operation func(a float64, b float64) float64, a float64, b float64) float64 {
	return operation(a, b)
}

func sum() {
	panic("Implement me")
}
func product() {
	panic("Implement me")
}
func divide() {
	panic("Implement me")
}

func main() {
	a := float64(2)
	b := float64(4)

	c := doOperation(sum, a, b)
	c = doOperation(product, c, a)
	c = doOperation(divide, c, b)

	if c != 3 {
		panic("Wrong")
	} else {
		fmt.Println("Correct ✅")
	}
}
