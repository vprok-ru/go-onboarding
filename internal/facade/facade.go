package facade

import (
	"main/internal/facade/after-hours"
	"main/internal/facade/food"
	"main/internal/facade/job"
	"main/internal/facade/sleep"
	"strings"
)

// Interface dayPlan
type Interface interface {
	Implement() string
}

// dayPlan Plan on day
type dayPlan struct {
	food       food.Interface
	job        job.Interface
	afterHours after_hours.Interface
	sleep      sleep.Interface
}

// Implement implementation of the plan for the day
func (p *dayPlan) Implement() string {
	result := [...]string{
		p.food.Eat(),
		p.job.Work(),
		p.afterHours.Relax(),
		p.sleep.Do(),
	}
	return strings.Join(result[:], "\n")
}

// NewDayPlan dayPlan fabric
func NewDayPlan(
	food food.Interface,
	job job.Interface,
	afterHours after_hours.Interface,
	sleep sleep.Interface,
) Interface {
	return &dayPlan{
		food,
		job,
		afterHours,
		sleep,
	}
}
