package after_hours

// Interface ...
type Interface interface {
	Relax() string
}

// afterHours struct
type afterHours struct {
}

// Relax ...
func (r *afterHours) Relax() string {
	return "Relax after hours"
}

// NewAfterHours fabric for afterHours with Interface
func NewAfterHours() Interface {
	return &afterHours{}
}
