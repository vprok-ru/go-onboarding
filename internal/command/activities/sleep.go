package activities

// sleep struct
type sleep struct {
}

// Do ...
func (s *sleep) Do() string {
	return "Sleep"
}

// NewSleep fabric for sleep with Activity interface
func NewSleep() Activity {
	return &sleep{}
}
