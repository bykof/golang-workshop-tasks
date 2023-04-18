package main

import "fmt"

/*
	Make the code compile
*/

func main() {
	a := int8(123)
	b := int16(321)
	fmt.Println(a + b)

	c := float32(3.14)
	d := float64(1.34)

	fmt.Println(c * d)

	e := int64(1000)
	f := float64(1.1)

	fmt.Println(e - f)
}
