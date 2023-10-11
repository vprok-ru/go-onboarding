package cars

// kia params
type kia struct {
	label  string
	weight uint
	power  uint
}

// SetLabel set car label
func (k *kia) SetLabel() {
	k.label = "Kia"
}

// SetWeight set car weight
func (k *kia) SetWeight() {
	k.weight = 1500
}

// SetPower set car power
func (k *kia) SetPower() {
	k.power = 150
}

// NewCar fabric for kia
func (k *kia) NewCar() Car {
	return NewCar(k.label, k.weight, k.power)
}

// NewKia fabric for kia with builder interface
func NewKia() Builder {
	return &kia{}
}
