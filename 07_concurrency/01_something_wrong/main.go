package main

import (
	"fmt"
	"time"
)

/*
TODO: There is a deadlock in this code, can you find out why?
*/

func calculate(done chan<- bool) {
	time.Sleep(3 * time.Second)
	done <- true
}

func main() {
	done := make(chan bool)
	go calculate(done)
	for range done {
		fmt.Println("Done")
	}

	fmt.Println("Exit")
}
