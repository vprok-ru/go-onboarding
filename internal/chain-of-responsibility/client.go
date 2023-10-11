package chain_of_responsibility

import "errors"

// carClient client handler
type carClient struct {
	next CarHandler
}

// CheckStatus check status
func (c *carClient) CheckStatus(status Status) (string, error) {
	if status == StatusClient {
		return "Client's car", nil
	} else if c.next != nil {
		return c.next.CheckStatus(status)
	}

	return "", errors.New("unknown status")
}

// NewCarClient fabric for carClient with CarHandler interface
func NewCarClient(next CarHandler) CarHandler {
	return &carClient{next}
}
