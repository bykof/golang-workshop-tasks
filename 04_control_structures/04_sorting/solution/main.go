package main

import (
	"fmt"
	"sort"
)

/*
   	Task:
       Sort Authors by different values

	Hint:
		Use "sort.Slice" from the sort package: https://pkg.go.dev/sort#Slice


   	Output:
    	[{Cina 634 102} {Mina 304 1098} {Mohit 300 104} {Rina 10 108} {Riya 4 101} {Rohit 56 107} {Sina 234 103} {Sohit 20 110} {Tina 104 105} {Vina 237 106}]
    	[{Riya 4 101} {Cina 634 102} {Sina 234 103} {Mohit 300 104} {Tina 104 105} {Vina 237 106} {Rohit 56 107} {Rina 10 108} {Sohit 20 110} {Mina 304 1098}]
    	[{Riya 4 101} {Rina 10 108} {Sohit 20 110} {Rohit 56 107} {Tina 104 105} {Sina 234 103} {Vina 237 106} {Mohit 300 104} {Mina 304 1098} {Cina 634 102}]



*/
func main() {
	authors := []struct {
		name     string
		articles int
		id       int
	}{
		{"Mina", 304, 1098},
		{"Cina", 634, 102},
		{"Tina", 104, 105},
		{"Rina", 10, 108},
		{"Sina", 234, 103},
		{"Vina", 237, 106},
		{"Rohit", 56, 107},
		{"Mohit", 300, 104},
		{"Riya", 4, 101},
		{"Sohit", 20, 110},
	}

	sort.Slice(authors, func(i int, j int) bool {
		return authors[i].name < authors[j].name
	})

	fmt.Println(authors)
	sort.Slice(authors, func(i int, j int) bool {
		return authors[i].id < authors[j].id
	})

	fmt.Println(authors)
	sort.Slice(authors, func(i int, j int) bool {
		return authors[i].articles < authors[j].articles
	})
	fmt.Println(authors)
}
