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

func (s Square) Area() float64 {
	return s.side * s.side
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func printArea(s Shape) {
	fmt.Printf("The area is %f\n", s.Area())
}

func main() {
	shapes := []Shape{
		Square{side: 10},
		Circle{radius: 10},
		Square{side: 20},
		Circle{radius: 8},
	}

	for _, v := range shapes {
		printArea(v)

		// Type Assertions
		var shape Shape = v

		_, circleOk := shape.(Circle)
		_, squareOk := shape.(Square)

		if circleOk {
			fmt.Println("The shape is circle!")
		} else if squareOk {
			fmt.Println("The shape is square!")
		} else {
			fmt.Println("Unknown shape!")
		}
		fmt.Println()
	}
}
