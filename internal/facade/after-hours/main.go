package after_hours

type Interface interface {
	Relax() string
}

type afterHours struct {
}

func (r *afterHours) Relax() string {
	return "Relax after hours"
}

func NewAfterHours() Interface {
	return &afterHours{}
}
