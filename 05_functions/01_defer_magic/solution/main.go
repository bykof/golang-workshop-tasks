package main

import (
	"bufio"
	"fmt"
	"os"
)

func PrintReverse(s string) {
	defer func() {
		fmt.Println()
	}()
	for _, r := range s {
		defer func(r rune) {
			fmt.Print(string(r))
		}(r)
	}
}

/*
   Task:
       Read a string from stdin.
       Implement a function called PrintReverse(s string),
	   which prints a string in the reverse order just by using defer!

   Output:
       > Hello
       olleH
       > World
       dlroW
*/
func main() {
	// Your code goes here
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		PrintReverse(text)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
