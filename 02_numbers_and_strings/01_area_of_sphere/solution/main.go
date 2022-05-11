package main

import (
	"fmt"
)

/*
   Task:
       Calculate the area of a sphere.
       Hint: The area of a sphere is: 4πr²

   Output:
       radius: 10 = area: 1256.64
*/
func main() {
	pi := 3.14159
	radius := 10
	var area float64
	area = 4 * pi * float64(radius*radius)
	// Your code goes here
	fmt.Printf("radius: %d = area: %.2f\n", radius, area)
}
