package chain_of_responsibility

import "errors"

type carFactory struct {
	next CarHandler
}

func (c *carFactory) CheckStatus(status Status) (string, error) {
	if status == StatusFactory {
		return "Car in factory", nil
	} else if c.next != nil {
		return c.next.CheckStatus(status)
	}

	return "", errors.New("unknown status")
}

func NewCarFactory(next CarHandler) CarHandler {
	return &carFactory{next}
}
