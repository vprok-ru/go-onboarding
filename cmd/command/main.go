package main

import (
	"fmt"
	"main/internal/command"
	"main/internal/command/after-hours"
	"main/internal/command/food"
	"main/internal/command/job"
	"main/internal/command/models"
	"main/internal/command/sleep"
)

func main() {
	activities := [...]models.Activity{food.NewFood(),
		job.NewJob(),
		after_hours.NewAfterHours(),
		sleep.NewSleep(),
	}

	dayPlan := command.NewDayPlan(activities[:])

	fmt.Println("Day plan:\n" + dayPlan.Implement())
}
