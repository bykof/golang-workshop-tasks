package main

import (
	"fmt"
	"time"
)

/*
   Task:
    Display a endless rotating tick

   Hint:
    Use the same line of console by printing "\r" at the start of a line
    Use time.Sleep(100 * time.Millisecond) to sleep to 100 milliseconds

   Output:
    \
    |
    /
    -
    \
   ...

*/
func main() {
	chars := "\\|/-"
	for {
		for _, r := range chars {
			fmt.Printf("\r%s", string(r))
			time.Sleep(100 * time.Millisecond)
		}
	}
}
