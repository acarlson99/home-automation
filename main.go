package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/acarlson99/home-automation/src/device"
	"github.com/acarlson99/home-automation/src/elgato"
	"github.com/acarlson99/home-automation/src/expr"
	"github.com/acarlson99/home-automation/src/schedule"

	hpb "github.com/acarlson99/home-automation/proto/go"
	tpb "google.golang.org/protobuf/encoding/prototext"
)

var (
	deviceConfigFile   string
	scheduleConfigFile string
)

func main() {
	flag.StringVar(&deviceConfigFile, "devices", "devices.textpb", "textproto config for proto/Device.proto")
	flag.StringVar(&scheduleConfigFile, "schedule", "schedule.textpb", "textproto config for proto/Automate.proto")
	flag.Parse()

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	textpb, err := os.ReadFile(deviceConfigFile)
	if err != nil {
		panic(err)
	}
	smartDevices := &hpb.Devices{}
	if err := tpb.Unmarshal(textpb, smartDevices); err != nil {
		panic(err)
	}

	textpb, err = os.ReadFile(scheduleConfigFile)
	if err != nil {
		panic(err)
	}
	events := hpb.Events{}
	if err := tpb.Unmarshal(textpb, &events); err != nil {
		panic(err)
	}
	fmt.Println(tpb.Format(&events))

	// register devices
	devices := []*device.Device{}
	for _, sdv := range smartDevices.GetDevice() {
		switch d := sdv.Device.(type) {
		case *hpb.SmartDevice_ElgatoLight:
			light := elgato.NewLight(d.ElgatoLight)
			newD := device.NewDevice(light)
			devices = append(devices, newD)
			err := device.RegisterDevice(newD)
			if err != nil {
				log.Fatal(err)
			}
		case *hpb.SmartDevice_GoveeLight:
		}
	}

	for _, e := range events.Event {
		conds := e.StartIf
		if conds == nil {
			continue
		}
		_, err := expr.EvalComparisons(conds)
		if err != nil {
			log.Fatalf("invalid start_if fails: %v\n", err)
		}
	}

	scheduler, err := schedule.DevicesEvents(devices, &events)
	if err != nil {
		log.Fatalf("Unexpected error creating devices: %v\n", err)
	}
	scheduler.Start()

	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()

	scheduler.Shutdown()
}
