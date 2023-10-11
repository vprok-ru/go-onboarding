package visitor

// Figure interface
type Figure interface {
	GetType() string
	Accept(v Visitor)
}
