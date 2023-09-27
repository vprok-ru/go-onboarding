package after_hours

import "main/internal/command/models"

type afterHours struct {
}

func (r *afterHours) Do() string {
	return "Relax after hours"
}

func NewAfterHours() models.Activity {
	return &afterHours{}
}
