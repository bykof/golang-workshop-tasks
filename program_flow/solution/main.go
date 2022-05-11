package main

import (
	"fmt"
	"os"
	"strings"
)

/*
   Task:
       Use strings.ToLower and string.ToUpper with a switch and arguments

   Output:
       go build main.go

       ./main [command] [input]

       ./main upper hello
       -> HELLO

       ./main lower HellO
       -> hello
*/

const (
	Upper = "upper"
	Lower = "lower"
)

func main() {
	// Your code goes here
	command := os.Args[1]
	input := os.Args[2]

	switch command {
	case Upper:
		fmt.Println(strings.ToUpper(input))
	case Lower:
		fmt.Println(strings.ToLower(input))
	default:
		fmt.Println("Please use [command] as lower or upper")
	}
}
