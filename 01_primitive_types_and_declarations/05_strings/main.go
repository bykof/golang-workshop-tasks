package main

import "fmt"

/*
Make the code compile and run successfully
*/

func main() {
	a := '®'
	var b string

	b = a
	fmt.Println(b)

	c := `HelloWorld`
	d := "Hello\nWorld"
	if c != d {
		panic("Wrong")
	} else {
		fmt.Println("Correct ✅")
	}
}
