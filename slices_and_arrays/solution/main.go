package main

import "fmt"

/*
   Task: Create an int array and a string slice.

   Int array is:
       - sizes[0] = 180
       - sizes[1] = 174
       - sizes[2] = 167
       - sizes[3] = 162

   String array is:
       - shirtSizes[0] = "XS"
       - shirtSizes[1] = "S"
       - shirtSizes[2] = "M"
       - shirtSizes[3] = "L"
       - shirtSizes[4] = "XL"

   Print out the the int array with index 2 and the string array with index 3

   Output:
       167 L
*/
func main() {
	// Your code goes here
	sizes := [4]int{
		180, 174, 167, 162,
	}
	shirtSizes := []string{
		"XS", "S", "M", "L", "XL",
	}

	fmt.Println(sizes[2], shirtSizes[3])
}
