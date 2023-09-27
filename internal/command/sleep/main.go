package sleep

import "main/internal/command/models"

type sleep struct {
}

func (s *sleep) Do() string {
	return "Sleep"
}

func NewSleep() models.Activity {
	return &sleep{}
}
