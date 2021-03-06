package timewheel_test

import (
	"algorithm/timewheel"
	"fmt"
	"time"
)

type EveryScheduler struct {
	Interval time.Duration
}

func (s *EveryScheduler)Next(prev time.Time) time.Time {
	return prev.Add(s.Interval)
}

func Example_scheduleTimer()  {
	tw := timewheel.NewTimingWheel(time.Millisecond,20)
	tw.Start()
	defer tw.Stop()
	exitc := make(chan time.Time)
	t := tw.ScheduleFunc(&EveryScheduler{time.Second}, func() {
		fmt.Println("The timer fires")
		exitc <- time.Now()
	})
	<-exitc
	<-exitc
	for !t.Stop() {

	}
}