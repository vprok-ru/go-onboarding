package chain_of_responsibility

import "errors"

// carFactory factory handler
type carFactory struct {
	next CarHandler
}

// CheckStatus check status
func (c *carFactory) CheckStatus(status Status) (string, error) {
	if status == StatusFactory {
		return "Car in factory", nil
	} else if c.next != nil {
		return c.next.CheckStatus(status)
	}

	return "", errors.New("unknown status")
}

// NewCarFactory fabric for carFactory with CarHandler interface
func NewCarFactory(next CarHandler) CarHandler {
	return &carFactory{next}
}
