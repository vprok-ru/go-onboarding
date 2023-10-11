package activities

// afterHours struct
type afterHours struct {
}

// Do ...
func (r *afterHours) Do() string {
	return "Relax after hours"
}

// NewAfterHours fabric for afterHours with Activity interface
func NewAfterHours() Activity {
	return &afterHours{}
}
