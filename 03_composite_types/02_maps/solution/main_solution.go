package main

import "fmt"

/*
   Task: Create a map and save the following values:

       - The Batman: 2022
       - Dune: 2021
       - The Godfather: 1972
       - Harry Potter and the Philisopher's Stone: 2002

   Overwrite "Harry Potter and the Philisopher's Stone" with the correct year 2001
   Print out all values.

   Output:
       - The Batman: 2022
       - Dune: 2021
       - The Godfather: 1972
       - Harry Potter and the Philisopher's Stone: 2001

*/
func main() {
	// Your code goes here
	movies := map[string]int{
		"The Batman":    2022,
		"Dune":          2021,
		"The Godfather": 1972,
		"Harry Potter and the Philisopher's Stone": 2002,
	}

	movies["Harry Potter and the Philisopher's Stone"] = 2001

	for key, value := range movies {
		fmt.Printf("- %s: %d\n", key, value)
	}
}
