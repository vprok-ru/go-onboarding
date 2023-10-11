package chain_of_responsibility

// Status status handler
type Status int

const (
	StatusFactory Status = iota
	StatusDealership
	StatusClient
	UnknownStatus
)
