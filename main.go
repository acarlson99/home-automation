package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/acarlson99/home-automation/src/elgato"

	hpb "github.com/acarlson99/home-automation/proto/go"
	tpb "google.golang.org/protobuf/encoding/prototext"
)

func main() {
	dname := flag.String("device-name", "", "device to modify")
	temperature := flag.Int("temperature", 150, "light temperature (if supported)")
	brightness := flag.Int("brightness", 50, "light brightness 0-100 (if supported)")
	on := flag.Int("on", 1, "turn on or off")
	flag.Parse()

	textpb, err := os.ReadFile("lights.textpb")
	if err != nil {
		panic(err)
	}
	smartDevices := &hpb.Devices{}
	if err := tpb.Unmarshal(textpb, smartDevices); err != nil {
		panic(err)
	}

	textpb, err = os.ReadFile("schedule.textpb")
	if err != nil {
		panic(err)
	}
	schedule := hpb.Events{}
	if err := tpb.Unmarshal(textpb, &schedule); err != nil {
		panic(err)
	}
	fmt.Println(tpb.Format(&schedule))

	// fmt.Println(smartDevices)
	for _, device := range smartDevices.GetDevice() {
		switch d := device.Device.(type) {
		case *hpb.SmartDevice_ElgatoLight:
			if d.ElgatoLight.Name != *dname {
				continue
			}
			light := elgato.NewLight(d.ElgatoLight)
			state := elgato.LightState{
				On:          *on,
				Temperature: *temperature,
				Brightness:  *brightness,
			}
			fmt.Printf("sending %+v\n", state)
			res, _ := light.SetLightVals(&elgato.LightsConfig{
				NumberOfLights: 1,
				Lights:         []elgato.LightState{state},
			})
			fmt.Printf("current state %+v\n", res)
		case *hpb.SmartDevice_GoveeLight:
		}
	}
}
