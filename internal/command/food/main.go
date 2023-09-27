package food

import "main/internal/command/models"

type food struct {
}

func (f *food) Do() string {
	return "Eat food"
}

func NewFood() models.Activity {
	return &food{}
}
