package food

type Interface interface {
	Eat() string
}

type food struct {
}

func (f *food) Eat() string {
	return "Eat food"
}

func NewFood() Interface {
	return &food{}
}
