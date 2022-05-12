package main

import (
	"fmt"
	"math"
)

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

type Circle struct {
	Radius float64
}

type Rectangle struct {
	A float64
	B float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.A * r.B
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.A + r.B)
}

var _ Shaper = Circle{}
var _ Shaper = Rectangle{}

func main() {
	// Your code goes here
	circle := Circle{
		Radius: 3,
	}
	rectangle := Rectangle{
		A: 123,
		B: 456,
	}
	fmt.Println(rectangle.Area())
	fmt.Println(rectangle.Perimeter())

	fmt.Println(circle.Area())
	fmt.Println(circle.Perimeter())
}
