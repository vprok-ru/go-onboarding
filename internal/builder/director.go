package builder

import (
	"main/internal/builder/cars"
)

// Interface for build car
type Interface interface {
	BuildCar() cars.Car
}

// director for creating cars
type director struct {
	builder cars.Builder
}

// BuildCar build car
func (d *director) BuildCar() cars.Car {
	d.builder.SetLabel()
	d.builder.SetWeight()
	d.builder.SetPower()
	return d.builder.NewCar()
}

// NewDirector fabric
func NewDirector(b cars.Builder) Interface {
	return &director{
		builder: b,
	}
}
