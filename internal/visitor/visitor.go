package visitor

// Visitor interface
type Visitor interface {
	VisitSquare(s *square)
	VisitTriangle(t *triangle)

	GetPerimeter() uint
}

// perimeter additional struct
type perimeter struct {
	p uint
}

// VisitTriangle for triangle
func (p *perimeter) VisitTriangle(t *triangle) {
	p.p = t.a + t.b + t.c
}

// VisitSquare for square
func (p *perimeter) VisitSquare(s *square) {
	p.p = s.a + s.b
}

// GetPerimeter get calculated perimeter
func (p *perimeter) GetPerimeter() uint {
	return p.p
}

// NewPerimeter fabric
func NewPerimeter() Visitor {
	return &perimeter{}
}
