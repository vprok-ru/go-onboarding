package convention

const (
	another = "qwwee"
)

var (
	some = "112"
)

type Iron struct {
	A int
	B string
}

type nakovalnya interface {
	Do() error
}

type molot interface {
	Do2() error
}

// Knife ...
type Knife struct {
	A int
	C float32
}

// Ker do some
type Ker interface {
	DoKnife(i Iron) (k Knife, err error)
}

type ker struct {
	nakovalnya nakovalnya
	molot molot
}

func (k *ker) DoKnife(i Iron) (kn Knife, err error) {
	_ = k.molot.Do2()
	err = k.nakovalnya.Do()
	kn.A = i.A
	return
}

// NewKer ...
func NewKer(nakovalnya nakovalnya, molot molot) Ker {
	return &ker{
		nakovalnya: nakovalnya,
		molot: molot,
	}
}
