package main

import (
	"flag"
	"log"
	"os"
	"sync"

	"github.com/acarlson99/home-automation/src/common"
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
	logLvl             string
)

func main() {
	flag.StringVar(&deviceConfigFile, "devices", "devices.textpb", "textproto config for proto/Device.proto")
	flag.StringVar(&scheduleConfigFile, "schedule", "schedule.textpb", "textproto config for proto/Automate.proto")
	flag.StringVar(&logLvl, "log-level", "warn", "level of detail to log: oneof error,warn,info,debug")
	flag.Parse()

	switch logLvl {
	case "error":
		common.SetLogLevel(common.Error)
	case "warn":
		common.SetLogLevel(common.Warn)
	case "info":
		common.SetLogLevel(common.Info)
	case "debug":
		common.SetLogLevel(common.Debug)
	default:
		log.Println("invalid log level:", logLvl)
		flag.Usage()
		os.Exit(1)
	}

	textpb, err := os.ReadFile(deviceConfigFile)
	if err != nil {
		common.Logger(common.Error).Printf("Error reading file %s: %v", deviceConfigFile, err)
		os.Exit(1)
	}
	smartDevices := hpb.Devices{}
	if err := tpb.Unmarshal(textpb, &smartDevices); err != nil {
		common.Logger(common.Error).Printf("Error unmarshalling proto: %v", err)
		os.Exit(1)
	}
	common.Logger(common.Info).Printf("%s:\n%s", deviceConfigFile, tpb.Format(&smartDevices))

	textpb, err = os.ReadFile(scheduleConfigFile)
	if err != nil {
		common.Logger(common.Error).Printf("Error reading file %s: %v", scheduleConfigFile, err)
		os.Exit(1)
	}
	events := hpb.Events{}
	if err := tpb.Unmarshal(textpb, &events); err != nil {
		common.Logger(common.Error).Printf("Error unmarshalling proto: %v", err)
		os.Exit(1)
	}
	common.Logger(common.Info).Printf("%s:\n%s", scheduleConfigFile, tpb.Format(&events))

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
				common.Logger(common.Error).Fatal(err)
			}
		case *hpb.SmartDevice_GoveeLight:
			common.Logger(common.Error).Fatal("unimplemented device")
		}
	}

	for _, e := range events.Event {
		conds := e.StartIf
		if conds == nil {
			continue
		}
		_, err := expr.EvalComparisons(conds)
		if err != nil {
			common.Logger(common.Error).Fatalf("invalid start_if fails: %v\n", err)
		}
	}

	scheduler, err := schedule.DevicesEvents(devices, &events)
	if err != nil {
		common.Logger(common.Error).Fatalf("Unexpected error creating devices: %v\n", err)
	}
	scheduler.Start()

	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()

	scheduler.Shutdown()
}
