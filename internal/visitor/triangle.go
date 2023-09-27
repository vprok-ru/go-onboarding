package visitor

type triangle struct {
	a, b, c uint
}

// Accept add visitor
func (t *triangle) Accept(v Visitor) {
	v.VisitTriangle(t)
}

func (t *triangle) GetType() string {
	return "triangle"
}

func NewTriangle(a, b, c uint) Figure {
	return &triangle{
		a: a,
		b: b,
		c: c,
	}
}
