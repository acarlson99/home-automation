package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
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
	port               int
	hostname           string
)

func main() {
	flag.StringVar(&deviceConfigFile, "devices", "devices.textpb", "textproto config for proto/Device.proto")
	flag.StringVar(&scheduleConfigFile, "schedule", "schedule.textpb", "textproto config for proto/Automate.proto")
	flag.StringVar(&logLvl, "log-level", "warn", "level of detail to log: oneof error,warn,info,debug")
	flag.IntVar(&port, "port", 8080, "port on which to serve web UI")
	flag.StringVar(&hostname, "hostname", "localhost", "host address")
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
			common.Logger(common.Error).Fatal("unimplemented device type SmartDevice_GoveeLight")
		}
	}

	for _, e := range events.Event {
		conds := e.StartIf
		if conds == nil {
			continue
		}
		_, err := expr.EvalComparisons(conds)
		if err != nil {
			common.Logger(common.Error).Fatalf("invalid start_if fails for device %v: %v\n", e.GetName(), err)
		}
	}

	scheduler, eventDevices, err := schedule.DevicesEvents(devices, &events)
	if err != nil {
		common.Logger(common.Error).Fatalf("Unexpected error creating devices: %v\n", err)
	}
	go spinup(hostname, port, eventDevices)
	scheduler.Start()

	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()

	scheduler.Shutdown()
}

func spinup(host string, port int, eds []*schedule.EventDevice) {
	buttonF := func(basePath, s string) string {
		return fmt.Sprintf(`<form action="%s/%s" method="post"> <button type="submit">%s</button> </form>`, basePath, s, s)
	}

	cmdURLBase := "/cmd"
	html := `
	<html>
	<head>
		<title>Button Clicker</title>
	</head>
	<body>
		<h1>Click the button:</h1>
		%s
	</body>
	</html>
	`
	bs := []string{}
	for _, ed_ := range eds {
		ed := ed_ // for scope
		event := ed.Event
		devs := []string{}
		for _, d := range ed.Ds {
			devs = append(devs, d.GetName())
		}
		actName := event.GetName() + "-" + strings.Join(devs, " ")
		http.HandleFunc(cmdURLBase+"/"+actName, func(w http.ResponseWriter, r *http.Request) {
			common.Logger(common.Debug).Printf("Running action from web-UI: %s", actName)
			fmt.Fprintf(w, html, strings.Join(bs, ""))
			go ed.RunEvent()
		})
		bs = append(bs, buttonF(cmdURLBase, actName))
	}
	// html = fmt.Sprintf(html, strings.Join(bs, ""))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, html, strings.Join(bs, ""))
	})

	common.Logger(common.Info).Printf("Server is running at http://%s:%v\n", host, port)
	go http.ListenAndServe(fmt.Sprintf("%s:%v", host, port), nil)
}
