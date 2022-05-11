package main

import "fmt"

/*
   Task:
       Swap the values of the variables in one line

   Output:
       Die Hard Harry Potter
*/
func main() {
	movie := "Harry Potter"
	movie2 := "Die Hard"
	movie, movie2 = movie2, movie
	fmt.Println(movie, movie2)
}
