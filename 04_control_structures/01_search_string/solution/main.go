package main

import (
	"fmt"
	"os"
	"strings"
)

/*
   Task:
       Pass an argument into your application and search for a movie (case insensitive, substring):

       +------------------------------------------+------+-------+
       |                   name                   | year | score |
       +------------------------------------------+------+-------+
       | The Batman                               | 2022 |     8 |
       | Dune                                     | 2021 |   8.1 |
       | The Godfather                            | 1972 |   9.2 |
       | Harry Potter and the Philosopher's Stone | 2001 |   7.6 |
       +------------------------------------------+------+-------+

   Hint:
       Use the "strings" library: https://pkg.go.dev/strings
       Use the "os" library with os.Args variable (it's a slice): https://pkg.go.dev/os

   Output:
       ./main ba
       will output:
           - The Batman(2022) Score: 8

       ./main a
       will output:
           - The Batman(2022) Score: 8
           - The Godfather(1972) Score: 9.2
           - Harry Potter and the Philosopher's Stone(2001) Score: 7.6

*/
func main() {
	// Your code goes here
	movies := []struct {
		name  string
		year  int64
		score float64
	}{
		{name: "The Batman", year: 2022, score: 8},
		{name: "The Godfather", year: 1972, score: 9.2},
		{name: "Dune", year: 2021, score: 8.1},
	}
	search := os.Args[1]

	for _, movie := range movies {
		if strings.Contains(strings.ToLower(movie.name), strings.ToLower(search)) {
			fmt.Printf(
				"- %s(%d) Score: %.1f\n",
				movie.name,
				movie.year,
				movie.score,
			)
		}
	}
}
