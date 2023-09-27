package cars

import "main/internal/builder/models"

type Builder interface {
	SetLabel()
	SetWeight()
	SetPower()
	NewCar() models.Car
}
