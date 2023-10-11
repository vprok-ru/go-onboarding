package cars

// Builder car interface
type Builder interface {
	SetLabel()
	SetWeight()
	SetPower()
	NewCar() Car
}
