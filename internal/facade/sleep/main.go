package sleep

type Interface interface {
	Do() string
}

type sleep struct {
}

func (s *sleep) Do() string {
	return "Sleep"
}

func NewSleep() Interface {
	return &sleep{}
}
