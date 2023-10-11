package strategy

// Strategy interface
type Strategy interface {
	search(array []string, value string) bool
}
