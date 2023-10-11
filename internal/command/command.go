package command

import (
	"strings"

	"main/internal/command/activities"
)

// Interface dayPlan
type Interface interface {
	Implement() string
}

// dayPlan plan on day
type dayPlan struct {
	activities []activities.Activity
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
	activities []activities.Activity,
) Interface {
	return &dayPlan{
		activities,
	}
}
