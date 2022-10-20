package main

import "fmt"

func main() {
	var x [4]int

	// Set the variable x so that the condition becomes true

	if x != [4]int{
		4, 3, 2, 1,
	} {
		panic("Wrong")
	} else {
		fmt.Println("Correct ✅")
	}
}
