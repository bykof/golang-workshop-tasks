package main

import (
	"fmt"
	"reflect"
)

/*
Make the code work
Don't care about reflect.DeepEqual for now, this only checks, if two slices are equal
*/

func main() {
	var x []string

	if x[0] != "Hello" && x[1] != "World" {
		panic("Wrong")
	} else {
		fmt.Println("Correct ✅")
	}

	var y = make([]int, 10)
	var z = []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}
	for index, _ := range y {
		y[index] = index
	}

	if !reflect.DeepEqual(y, z) {
		panic("Wrong")
	} else {
		fmt.Println("Correct ✅")
	}

	if !reflect.DeepEqual(y[3:9], []int{
		4, 5, 6,
	}) {
		panic("Wrong")
	} else {
		fmt.Println("Correct ✅")
	}

}
