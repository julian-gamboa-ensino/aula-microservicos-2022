package main

import (
	"agosto-24/circle"
	"agosto-24/square"
	"fmt"
)

type Shape interface {
	Area() float64
}

// A helper function to calculate the area of any Shape
func calculateArea(s Shape) float64 {
	return s.Area()
}

func main() {

	fmt.Println(calculateArea(circle.Circle{Radius: 10}))

	fmt.Println(calculateArea(square.Square{Side: 10}))
}
