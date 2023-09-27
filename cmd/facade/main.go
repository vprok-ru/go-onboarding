package main

import (
	"fmt"
	"main/internal/facade"
	"main/internal/facade/after-hours"
	"main/internal/facade/food"
	"main/internal/facade/job"
	"main/internal/facade/sleep"
)

func main() {
	dayPlan := facade.NewDayPlan(
		food.NewFood(),
		job.NewJob(),
		after_hours.NewAfterHours(),
		sleep.NewSleep(),
	)
	fmt.Println("Day plan:\n" + dayPlan.Implement())
}
