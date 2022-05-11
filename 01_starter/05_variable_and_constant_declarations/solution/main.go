package main

import "fmt"

/*
   Task:
       Declare and Initialize all possible variable and constant formats

   Output:
       Harry Potter Die Hard The Batman Dune The Godfather
*/
func main() {
	var movie string = "Harry Potter"
	var movie2 = "Die Hard"
	movie3 := "The Batman"
	movie3, movie4 := movie3, "Dune"
	const movie5 string = "The Godfather"
	const movie6 = "The Godfather"
	// Your code goes here
	fmt.Println(movie, movie2, movie3, movie4, movie5, movie6)
}
