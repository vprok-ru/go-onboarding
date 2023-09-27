package chain_of_responsibility

type CarHandler interface {
	CheckStatus(status Status) (string, error)
}
