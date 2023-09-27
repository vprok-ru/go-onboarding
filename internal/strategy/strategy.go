package strategy

type StrategyRelease interface {
	SetStrategy(strategy Strategy)
	Search(array []string, value string) bool
}

type strategyRelease struct {
	strategy Strategy
}

// SetStrategy update strategy
func (s *strategyRelease) SetStrategy(strategy Strategy) {
	s.strategy = strategy
}

// Search method call strategy.search(...)
func (s *strategyRelease) Search(array []string, value string) bool {
	return s.strategy.search(array, value)
}

// NewStrategyRelease strategyRelease fabric
func NewStrategyRelease() StrategyRelease {
	return &strategyRelease{}
}
