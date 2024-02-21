package controller

import (
	"log"

	"github.com/acarlson99/home-automation/src/common"
	"github.com/acarlson99/home-automation/src/device"
	"github.com/acarlson99/home-automation/src/expr"

	hpb "github.com/acarlson99/home-automation/proto/go"
)

func RunEvent(devices []*device.Device, event *hpb.Event) {
	dnames := []string{}
	for _, d := range devices {
		dnames = append(dnames, d.GetName())
	}
	log.Println("running scheduled routine", event.GetName(), "for devices", dnames)

	do, err := expr.EvalComparisons(event.GetStartIf())
	if err != nil {
		log.Printf("Error evaluating comparison: %v\n", err)
	}
	if !do {
		log.Println("comparison returned false, skipping")
		return
	}

	f := func(d *device.Device) error { return d.ExecuteAll(event.Actions) }
	err = common.ConcurrentAggregateErrorFn(f, devices...)
	if err != nil {
		// TODO: more advanced error reporting than this
		log.Printf("Error executing event %v: %v\n", event.Name, err)
	}
}
