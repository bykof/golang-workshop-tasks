package main

import (
	"fmt"
	"math/rand"
)

/*
	Tasks:
		Create n goroutines to start an endless for loop.
		Each goroutine should get a randomized rand.Intn(2000) number.
		If the number becomes 1337 stop all goroutines and print which goroutine (numerized) won.

	Output:
		And the winner is 12
*/

func main() {
	n := 10
	won := make(chan int)
	for i := 0; i < n; i++ {
		go func(i int, won chan<- int) {
			for {
				number := rand.Intn(2000000)
				if number == 1333337 {
					won <- i
				}
			}
		}(i, won)
	}
	result := <-won
	fmt.Printf("And the winner is: %d", result)
}
