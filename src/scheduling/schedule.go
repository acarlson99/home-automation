package scheduling

import (
	"fmt"
	"log"

	"github.com/go-co-op/gocron/v2"

	hpb "github.com/acarlson99/home-automation/proto/go"
	device_controller "github.com/acarlson99/home-automation/src/device-controller"
)

func eventTimes(event *hpb.Event) []gocron.AtTime {
	ts := []gocron.AtTime{}
	for _, sched := range event.GetSchedule() {
		at := sched.GetAt() // TODO: adapt for different scheduling mechanisms
		ts = append(ts, gocron.NewAtTime(uint(at.Hour), uint(at.Minute), uint(at.Second)))
	}
	return ts
}

func sched(events *hpb.Events) (cleanup func() error, err error) {
	// create a scheduler
	s, err := gocron.NewScheduler()
	if err != nil {
		// handle error
	}

	// add a job to the scheduler
	for _, event := range events.Event {
		ts := eventTimes(event)
		var times gocron.AtTimes
		if len(ts) > 1 {
			times = gocron.NewAtTimes(ts[0], ts[1:]...)
		} else {
			times = gocron.NewAtTimes(ts[0])
		}

		j, err := s.NewJob(
			gocron.DailyJob(1, times),
			// gocron.DurationJob(
			// 	10*time.Second,
			// ),
			gocron.NewTask(
				func() {
					event := event
					// TODO: select which device to associate with this
					var d *device_controller.Device
					err := d.ExecuteAll(event.Actions)
					if err != nil {
						log.Printf("error executing %v: %v", event.Name, err)
					}
				},
			),
		)
		if err != nil {
			// handle error
		}
		// each job has a unique id
		fmt.Println(j.ID())
	}
	// start the scheduler
	s.Start()

	// when you're done, shut it down
	return s.Shutdown, nil
}
