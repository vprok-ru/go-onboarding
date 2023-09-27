package strategy

type Strategy interface {
	search(array []string, value string) bool
}
