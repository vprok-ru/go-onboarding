package food

// Interface ...
type Interface interface {
	Eat() string
}

// food struct
type food struct {
}

// Eat ...
func (f *food) Eat() string {
	return "Eat food"
}

// NewFood fabric for food with Interface
func NewFood() Interface {
	return &food{}
}
