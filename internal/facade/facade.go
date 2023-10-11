package facade

import "strings"

// Sleep interface
type Sleep interface {
	Do() string
}

// AfterHours interface
type AfterHours interface {
	Relax() string
}

// Job interface
type Job interface {
	Work() string
}

// Food interface
type Food interface {
	Eat() string
}

// Interface dayPlan
type Interface interface {
	Implement() string
}

// dayPlan Plan on day
type dayPlan struct {
	food       Food
	job        Job
	afterHours AfterHours
	sleep      Sleep
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
	food Food,
	job Job,
	afterHours AfterHours,
	sleep Sleep,
) Interface {
	return &dayPlan{
		food,
		job,
		afterHours,
		sleep,
	}
}
