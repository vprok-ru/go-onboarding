package cars

import "main/internal/builder/models"

type kia struct {
	label  string
	weight uint
	power  uint
}

func (k *kia) SetLabel() {
	k.label = "Kia"
}

func (k *kia) SetWeight() {
	k.weight = 1500
}

func (k *kia) SetPower() {
	k.power = 150
}

func (k *kia) NewCar() models.Car {
	return models.NewCar(k.label, k.weight, k.power)
}

func NewKia() Builder {
	return &kia{}
}
