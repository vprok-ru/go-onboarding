package chain_of_responsibility

// CarHandler handler interface
type CarHandler interface {
	CheckStatus(status Status) (string, error)
}
