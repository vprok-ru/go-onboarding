package chain_of_responsibility

import "errors"

type carClient struct {
	next CarHandler
}

func (c *carClient) CheckStatus(status Status) (string, error) {
	if status == StatusClient {
		return "Client's car", nil
	} else if c.next != nil {
		return c.next.CheckStatus(status)
	}

	return "", errors.New("unknown status")
}

func NewCarClient(next CarHandler) CarHandler {
	return &carClient{next}
}
