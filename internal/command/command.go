package command

import (
	"strings"
)

// Activity interface
type Activity interface {
	Do() string
}

// DayPlan interface
type DayPlan interface {
	Implement() string
}

// dayPlan plan on day
type dayPlan struct {
	activities []Activity
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
	activities []Activity,
) DayPlan {
	return &dayPlan{
		activities,
	}
}
