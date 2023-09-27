package visitor

type Figure interface {
	GetType() string
	Accept(v Visitor)
}
