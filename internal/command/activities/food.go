package activities

// food struct
type food struct {
}

// Do ...
func (f *food) Do() string {
	return "Eat food"
}

// NewFood fabric for food with Activity interface
func NewFood() Activity {
	return &food{}
}
