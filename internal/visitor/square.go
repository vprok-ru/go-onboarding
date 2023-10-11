package visitor

// square has a, b params
type square struct {
	a, b uint
}

// Accept add visitor
func (s *square) Accept(v Visitor) {
	v.VisitSquare(s)
}

// GetType ...
func (s *square) GetType() string {
	return "square"
}

// NewSquare fabric for square with Figure interface
func NewSquare(a, b uint) Figure {
	return &square{
		a: a,
		b: b,
	}
}
