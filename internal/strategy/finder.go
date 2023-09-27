package strategy

type finder struct {
}

func (f *finder) search(array []string, value string) bool {
	for _, elem := range array {
		if elem == value {
			return true
		}
	}
	return false
}

func NewFinder() Strategy {
	return &finder{}
}
