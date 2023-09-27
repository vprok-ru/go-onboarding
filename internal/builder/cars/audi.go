package cars

import "main/internal/builder/models"

type audi struct {
	label  string
	weight uint
	power  uint
}

func (a *audi) SetLabel() {
	a.label = "Audi"
}

func (a *audi) SetWeight() {
	a.weight = 2000
}

func (a *audi) SetPower() {
	a.power = 300
}

func (a *audi) NewCar() models.Car {
	return models.NewCar(a.label, a.weight, a.power)
}

func NewAudi() Builder {
	return &audi{}
}
