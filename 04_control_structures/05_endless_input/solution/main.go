package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
   Task:
       Make a program where you input a line and the line gets printed inverted

   Hint:
       https://pkg.go.dev/bufio#example-Scanner-Lines

   Output:
       > Test
       tseT
       > Hello World
       dlroW olleH



*/
func main() {
	// Your code goes here
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := []byte(scanner.Text())
		for i, j := 0, len(text)-1; i < j; i, j = i+1, j-1 {
			text[i], text[j] = text[j], text[i]
		}

		fmt.Println(string(text))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
