package main

import (
	"fmt"

	"main/internal/builder"
	"main/internal/builder/cars"
)

func main() {
	audi := cars.NewAudi()
	factory := builder.NewDirector(audi)
	car := factory.BuildCar()

	fmt.Println(car.GetParams() + "\n")

	kia := cars.NewKia()
	factory = builder.NewDirector(kia)
	car = factory.BuildCar()

	fmt.Println(car.GetParams() + "\n")
}
