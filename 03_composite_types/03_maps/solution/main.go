package main

import "fmt"

/*
   Task:
       Access the map with an existing and non existing key using the variable, ok := idiom
       Print out the variables.

   Output:
       29 true
       0 false
*/
func main() {
	ages := map[string]int{
		"Sarah": 29,
		"Tom":   24,
	}

	sarah, ok := ages["Sarah"]
	fmt.Println(sarah, ok)
	notExisting, ok := ages["NotExisting"]
	fmt.Println(notExisting, ok)
	// Your code goes here
}
