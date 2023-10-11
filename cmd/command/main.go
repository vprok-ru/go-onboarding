package main

import (
	"fmt"

	"main/internal/command"
	"main/internal/command/activities"
)

func main() {
	activitiesArr := [...]command.Activity{activities.NewFood(),
		activities.NewJob(),
		activities.NewAfterHours(),
		activities.NewSleep(),
	}

	dayPlan := command.NewDayPlan(activitiesArr[:])

	fmt.Println("Day plan:\n" + dayPlan.Implement())
}
