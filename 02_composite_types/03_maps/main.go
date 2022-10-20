package main

import "fmt"

/*
Make code compile and work
*/

func main() {
	var people []map[string]string

	if people[0]["firstName"] != "Michael" && people[0]["lastName"] != "Tester" {
		panic("Wrong")
	} else {
		fmt.Println("Correct ✅")
	}

	if people[1]["firstName"] != "Chris" && people[1]["lastName"] != "Martin" {
		panic("Wrong")
	} else {
		fmt.Println("Correct ✅")
	}

	// Use the var, ok idiom to check if the key "nope" exists on the first map element
	fmt.Println(firstElement, ok)
	if ok {
		panic("Wronf")
	} else {
		fmt.Println(firstElement)
		fmt.Println("Correct ✅")
	}
}
