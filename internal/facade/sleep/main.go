package sleep

// Interface ...
type Interface interface {
	Do() string
}

// sleep struct
type sleep struct {
}

// Do ...
func (s *sleep) Do() string {
	return "Sleep"
}

// NewSleep fabric for sleep with Interface
func NewSleep() Interface {
	return &sleep{}
}
