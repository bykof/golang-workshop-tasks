package main

/*
   Task:
       Implement two structs which implement the Shaper interface: Rectangle and Circle
       Print out some examples.

   Hint:
       Area of a circle:
           PI * radius * radius
       Perimeter of a circle:
           2 * PI * radius

       Area of a rectangle:
           a * b
       Perimeter of an Area:
           2 * (a + b)
*/

type Shaper interface {
	Area() float64
	Perimeter() float64
}

func main() {
	// Your code goes here
}
