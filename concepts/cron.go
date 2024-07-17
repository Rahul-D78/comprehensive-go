package concepts

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func GetScheduleTime() {
	now := time.Now().UTC()
	sched, err := cron.ParseStandard("*/5 * * * *")
	if err != nil {
		fmt.Println(err.Error())
	}
	nextTriggerTimeA := sched.Next(now)
	endTimeA := nextTriggerTimeA.Add(5 * time.Minute)
	nextTriggerTimeB := sched.Next(nextTriggerTimeA)
	// requeueAfter := time.Until(nextTriggerTime)
	// fmt.Println(requeueAfter, "<<<<---- requeafter")
	fmt.Println(nextTriggerTimeA, "next schedule time A")
	fmt.Println(nextTriggerTimeB, "next schedule time B")
	fmt.Println(endTimeA, "end time of first event")
}
