package main

import (
	"agosto-24/Areas/circle"
	"agosto-24/Areas/square"
	"fmt"
)

type Shape2D interface {
	Area() float64
}

type Shape3D interface {
	Shape2D
	Volume() float64
}

// A helper function to calculate the area of any Shape
func calculateArea(s Shape2D) float64 {
	return s.Area()
}

func main() {

	fmt.Println(calculateArea(circle.Circle{Radius: 10}))

	fmt.Println(calculateArea(square.Square{Side: 10}))
}
