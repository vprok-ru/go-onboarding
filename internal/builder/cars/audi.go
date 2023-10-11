package cars

// audi params
type audi struct {
	label  string
	weight uint
	power  uint
}

// SetLabel set car label
func (a *audi) SetLabel() {
	a.label = "Audi"
}

// SetWeight set car weight
func (a *audi) SetWeight() {
	a.weight = 2000
}

// SetPower set car power
func (a *audi) SetPower() {
	a.power = 300
}

// NewCar fabric for audi
func (a *audi) NewCar() Car {
	return NewCar(a.label, a.weight, a.power)
}

// NewAudi fabric for audi with builder interface
func NewAudi() Builder {
	return &audi{}
}
