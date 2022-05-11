package main

import "fmt"

/*
   Task:
       Slice out the two "Go" in the text and print them out

   Output:
       Go
       Go
*/

func main() {
	text := "Wake me up before you Go Go"
	fmt.Println(text[22:24])
	fmt.Println(text[25:27])
	// Your code goes here
}
