package main

import "fmt"

/*
   Task: Create a struct slice and save the following values:

       +------------------------------------------+------+-------+
       |                   name                   | year | score |
       +------------------------------------------+------+-------+
       | The Batman                               | 2022 |     8 |
       | Dune                                     | 2021 |   8.1 |
       | The Godfather                            | 1972 |   9.2 |
       | Harry Potter and the Philosopher's Stone | 2001 |   7.6 |
       +------------------------------------------+------+-------+

   Output:
       - The Batman(2022) Score: 8.0
       - Dune(2021) Score: 8.1
       - The Godfather(1972) Score: 9.2
       - Harry Potter and the Philisopher's Stone(2001) Score: 7.6

*/
func main() {
	// Your code goes here
	movies := []struct {
		name  string
		year  int64
		score float64
	}{
		{name: "The Batman", year: 2022, score: 8},
		{name: "Dune", year: 2021, score: 8.1},
	}

	for _, movie := range movies {
		fmt.Printf(
			"- %s(%d) Score: %.1f\n",
			movie.name,
			movie.year,
			movie.score,
		)
	}
}
