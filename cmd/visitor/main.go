package main

import (
	"fmt"

	"main/internal/visitor"
)

func main() {
	triandle := visitor.NewTriangle(10, 5, 10)
	square := visitor.NewSquare(20, 10)

	perimeterCalculator := visitor.NewPerimeter()

	triandle.Accept(perimeterCalculator)
	fmt.Println(fmt.Sprintf("Triangle perimeter: %d", perimeterCalculator.GetPerimeter()))

	square.Accept(perimeterCalculator)
	fmt.Println(fmt.Sprintf("Square perimeter: %d", perimeterCalculator.GetPerimeter()))
}
