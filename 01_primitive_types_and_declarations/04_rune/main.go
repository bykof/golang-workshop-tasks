package main

import "fmt"

/*
	Calculate a so it becomes ® (rights reserved symbol)

	Hint: The solution a = '®' is wrong!
	Hint the decimal value of '±' is 49841 and '®' is 49838
*/

func main() {
	a := '±'
	if a == '®' {
		fmt.Printf("Yes you did it: %q\n", a)
	} else {
		panic("Nope it's not correct")
	}
}
