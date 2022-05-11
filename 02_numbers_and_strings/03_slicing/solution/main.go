package main

import "fmt"

/*
   Task:
       Take "hello" and "programming" and print out just the
       half (round to floor) of the word length by using slicing

   Output:
       he
       progr
*/

func main() {
	hello := "hello"
	programming := "programming"

	fmt.Println(hello[:len(hello)/2])
	fmt.Println(programming[:len(programming)/2])
	// Your code goes here
}
