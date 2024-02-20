package schedule

import (
	"fmt"
	"log"

	"github.com/acarlson99/home-automation/src/common"
	"github.com/acarlson99/home-automation/src/controller"
	"github.com/go-co-op/gocron/v2"

	hpb "github.com/acarlson99/home-automation/proto/go"
)

func eventTimes(event *hpb.Event) []gocron.AtTime {
	ts := []gocron.AtTime{}
	for _, sched := range event.GetSchedule() {
		at := sched.GetDaily() // TODO: adapt for different scheduling mechanisms
		ts = append(ts, gocron.NewAtTime(uint(at.GetHour()), uint(at.GetMinute()), uint(at.GetSecond())))
	}
	return ts
}

type Scheduler gocron.Scheduler

func DevicesEvents(devices []*controller.Device, events *hpb.Events) (Scheduler, error) {
	s, err := gocron.NewScheduler()
	if err != nil {
		return nil, fmt.Errorf("could not create gocron scheduler: %v", err)
	}

	for _, event := range events.Event {
		var jobDef gocron.JobDefinition

		if len(event.GetSchedule()) == 1 && event.GetSchedule()[0].GetCrontab() != "" {
			jobDef = gocron.CronJob(event.GetSchedule()[0].GetCrontab(), false)
		} else {
			ts := eventTimes(event)
			var times gocron.AtTimes
			if len(ts) > 1 {
				times = gocron.NewAtTimes(ts[0], ts[1:]...)
			} else {
				times = gocron.NewAtTimes(ts[0])
			}
			jobDef = gocron.DailyJob(1, times)
		}

		eventDevices := []*controller.Device{}
		for _, dev := range devices {
			for _, dev2 := range event.GetDevices() {
				if dev.NameMatches(dev2.GetName()) {
					eventDevices = append(eventDevices, dev)
				}
			}
		}

		j, err := s.NewJob(jobDef, gocron.NewTask(devicesRunEvent, eventDevices, event))
		if err != nil {
			return nil, fmt.Errorf("could not create gocron job: %v", err)
		}
		log.Println("scheduled id", j.ID(), "for", event.Name)
	}

	return s, nil
}

func devicesRunEvent(devices []*controller.Device, event *hpb.Event) {
	dnames := []string{}
	for _, d := range devices {
		dnames = append(dnames, d.GetName())
	}
	log.Println("running scheduled routine", event.GetName(), "for devices", dnames)

	f := func(d *controller.Device) error { return d.ExecuteAll(event.Actions) }
	err := common.ConcurrentAggregateErrorFn(f, devices...)
	if err != nil {
		// TODO: more advanced error reporting than this
		log.Printf("Error executing event %v: %v\n", event.Name, err)
	}
}
