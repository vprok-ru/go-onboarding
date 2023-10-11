package chain_of_responsibility

import "errors"

// carDealership dealership handler
type carDealership struct {
	next CarHandler
}

// CheckStatus check status
func (c *carDealership) CheckStatus(status Status) (string, error) {
	if status == StatusDealership {
		return "Сar in dealership", nil
	} else if c.next != nil {
		return c.next.CheckStatus(status)
	}

	return "", errors.New("unknown status")
}

// NewCarDealership fabric for carDealership with CarHandler interface
func NewCarDealership(next CarHandler) CarHandler {
	return &carDealership{next}
}
