package controller

import (
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
	common.Logger(common.Info).Printf("running scheduled routine \"%s\" for devices %v", event.GetName(), dnames)

	do, err := expr.EvalComparisons(event.GetStartIf())
	if err != nil {
		common.Logger(common.Error).Printf("Error evaluating comparison: %v\n", err)
	}
	if !do {
		common.Logger(common.Debug).Println("comparison returned false, skipping")
		return
	}

	f := func(d *device.Device) error { return d.ExecuteAll(event.Actions) }
	err = common.ConcurrentAggregateErrorFn(f, devices...)
	if err != nil {
		// TODO: more advanced error reporting than this
		common.Logger(common.Error).Printf("Error executing event %v: %v\n", event.Name, err)
	}
}
