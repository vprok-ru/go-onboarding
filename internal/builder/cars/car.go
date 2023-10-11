package cars

import (
	"fmt"
	"strings"
)

// Car interface
type Car interface {
	GetParams() string
}

// car params
type car struct {
	label  string
	weight uint
	power  uint
}

// GetParams return params of car in string
func (c *car) GetParams() string {
	res := []string{
		fmt.Sprintf("Model: %s", c.label),
		fmt.Sprintf("Weight: %d", c.weight),
		fmt.Sprintf("Power: %d", c.power),
	}
	return strings.Join(res, "\n")
}

// NewCar car fabric
func NewCar(label string, weight, power uint) Car {
	return &car{label, weight, power}
}
