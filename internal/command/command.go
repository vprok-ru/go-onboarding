package command

import (
	"main/internal/command/models"
	"strings"
)

// Interface dayPlan
type Interface interface {
	Implement() string
}

// dayPlan plan on day
type dayPlan struct {
	activities []models.Activity
}

// Implement implementation of the plan for the day
func (p *dayPlan) Implement() string {
	result := make([]string, len(p.activities))

	for i, activity := range p.activities {
		done := activity.Do()
		result[i] = done
	}

	return strings.Join(result, "\n")
}

// NewDayPlan dayPlan fabric
func NewDayPlan(
	activities []models.Activity,
) Interface {
	return &dayPlan{
		activities,
	}
}
