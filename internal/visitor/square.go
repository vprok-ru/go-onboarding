package visitor

type square struct {
	a, b uint
}

// Accept add visitor
func (s *square) Accept(v Visitor) {
	v.VisitSquare(s)
}

func (s *square) GetType() string {
	return "square"
}

func NewSquare(a, b uint) Figure {
	return &square{
		a: a,
		b: b,
	}
}
