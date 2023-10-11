package visitor

// triangle has a, b, c params
type triangle struct {
	a, b, c uint
}

// Accept add visitor
func (t *triangle) Accept(v Visitor) {
	v.VisitTriangle(t)
}

// GetType ...
func (t *triangle) GetType() string {
	return "triangle"
}

// NewTriangle fabric for triangle with Figure interface
func NewTriangle(a, b, c uint) Figure {
	return &triangle{
		a: a,
		b: b,
		c: c,
	}
}
