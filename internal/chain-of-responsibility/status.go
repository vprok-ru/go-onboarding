package chain_of_responsibility

type Status int

const (
	StatusFactory Status = iota
	StatusDealership
	StatusClient
	UnknownStatus
)
