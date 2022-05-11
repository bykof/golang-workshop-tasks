package main

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
}
