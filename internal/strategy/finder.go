package strategy

// finder struct
type finder struct {
}

// search method
func (f *finder) search(array []string, value string) bool {
	for _, elem := range array {
		if elem == value {
			return true
		}
	}
	return false
}

// NewFinder fabric for finder with Strategy interface
func NewFinder() Strategy {
	return &finder{}
}
