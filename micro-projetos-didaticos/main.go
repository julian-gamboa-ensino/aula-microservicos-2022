package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Square struct {
	side float64
}

func (sq Square) Area() float64 {
	return sq.side * sq.side
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

// A helper function to calculate the area of any Shape
func calculateArea(s Shape) float64 {
	return s.Area()
}

func main() {

	fmt.Println(calculateArea(Circle{radius: 10}))
}
