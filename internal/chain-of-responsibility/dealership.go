package chain_of_responsibility

import "errors"

type carDealership struct {
	next CarHandler
}

func (c *carDealership) CheckStatus(status Status) (string, error) {
	if status == StatusDealership {
		return "Сar in dealership", nil
	} else if c.next != nil {
		return c.next.CheckStatus(status)
	}

	return "", errors.New("unknown status")
}

func NewCarDealership(next CarHandler) CarHandler {
	return &carDealership{next}
}
