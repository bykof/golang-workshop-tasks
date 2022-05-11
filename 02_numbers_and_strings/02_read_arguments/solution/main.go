package main

import (
	"fmt"
	"os"
)

/*
   Task:
        Read first two arguments and print them
        Check: https://pkg.go.dev/os#pkg-variables

        os.Args[1] is the first argument
        os.Args[2] is the second argument

        Hint:
        Try first:
            go run main.go one two

        and then:
            go build main.go
            ./main one two


   Output:
       one two
*/
func main() {
	fmt.Println(os.Args[1], os.Args[2])
}
