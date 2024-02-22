package controller_test

import (
	"testing"

	hpb "github.com/acarlson99/home-automation/proto/go"
	"github.com/acarlson99/home-automation/src/controller"
	"github.com/acarlson99/home-automation/src/device"
	"github.com/acarlson99/home-automation/src/elgato"
	"github.com/acarlson99/home-automation/src/elgato/testutil"
	"github.com/google/go-cmp/cmp"
	tpb "google.golang.org/protobuf/encoding/prototext"
)

func TestRunEvent(t *testing.T) {
	// 0-100 makes it real easy to thing about percents
	elgato.MaxBrightness = 100
	elgato.MinBrightness = 0
	elgato.MaxColorTemperature = 100
	elgato.MinColorTemperature = 0
	conf := elgato.LightsConfig{
		NumberOfLights: 1,
		Lights: []elgato.LightState{
			{
				On:          1,
				Brightness:  50,
				Temperature: 20,
			},
		},
	}
	lname := "test-light"
	defaultDevice, cleanup := testutil.NewFakeLightServer(t, lname, &conf)
	defer cleanup()

	type args struct {
		devices []*device.Device
		event   string
	}
	tests := []struct {
		name string
		args args
		want *elgato.LightsConfig
	}{
		{
			name: "static set brightness",
			args: args{
				devices: []*device.Device{device.NewDevice(defaultDevice)},
				event: `name: "lights 50% warm"
devices {
	name: "test-light"
}
actions {
	brightness: 50
}
actions {
	color_temp: 100
}
schedule {
	daily {
		hour: 20
		minute: 0
		second: 0
	}
}`,
			},
			want: &elgato.LightsConfig{
				NumberOfLights: 1,
				Lights: []elgato.LightState{
					{
						On:          1,
						Brightness:  50,
						Temperature: 100,
					},
				},
			},
		},
		{
			name: "change relative brightness",
			args: args{
				devices: []*device.Device{device.NewDevice(defaultDevice)},
				event: `name: "temp bright -20%"
devices {
	name: "test-light"
}
actions {
	brightness: -20
	relative: true
}
actions {
	color_temp: -20
	relative: true
}
schedule {
	daily {
		hour: 20
		minute: 0
		second: 0
	}
}`,
			},
			want: &elgato.LightsConfig{
				NumberOfLights: 1,
				Lights: []elgato.LightState{
					{
						On:          1,
						Brightness:  50 - 20,
						Temperature: 100 - 20,
					},
				},
			},
		},
		{
			name: "change relative brightness",
			args: args{
				devices: []*device.Device{device.NewDevice(defaultDevice)},
				event: `name: "temp bright -20 again%"
devices {
	name: "test-light"
}
actions {
	brightness: -20
	relative: true
}
actions {
	color_temp: -20
	relative: true
}
schedule {
	daily {
		hour: 20
		minute: 0
		second: 0
	}
}`,
			},
			want: &elgato.LightsConfig{
				NumberOfLights: 1,
				Lights: []elgato.LightState{
					{
						On:          1,
						Brightness:  50 - 20*2,
						Temperature: 100 - 20*2,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			event := &hpb.Event{}
			tpb.Unmarshal([]byte(tt.args.event), event)
			controller.RunEvent(tt.args.devices, event)
		})
	}
}

func TestRunIterativeEvent(t *testing.T) {
	// 0-100 makes it real easy to thing about percents
	elgato.MaxBrightness = 100
	elgato.MinBrightness = 0
	elgato.MaxColorTemperature = 420
	elgato.MinColorTemperature = 69
	conf := elgato.LightsConfig{
		NumberOfLights: 1,
		Lights: []elgato.LightState{
			{
				On:          1,
				Brightness:  95,
				Temperature: 400,
			},
		},
	}
	lname := "test-light"
	deviceServer, cleanup := testutil.NewFakeLightServer(t, lname, &conf)
	defaultDevice := device.NewDevice(deviceServer)
	defer cleanup()

	s := `name: "temp bright -20 again%"
	devices {
		name: "test-light"
	}
	actions {
		brightness: -20
		relative: true
	}
	actions {
		color_temp: -10
		relative: true
	}
	schedule {
		daily {
			hour: 20
			minute: 0
			second: 0
		}
	}`
	event := &hpb.Event{}
	tpb.Unmarshal([]byte(s), event)

	tests := []*elgato.LightsConfig{
		&elgato.LightsConfig{
			NumberOfLights: 1,
			Lights: []elgato.LightState{
				{
					On:          1,
					Brightness:  75,
					Temperature: 365,
				},
			},
		},
		&elgato.LightsConfig{
			NumberOfLights: 1,
			Lights: []elgato.LightState{
				{
					On:          1,
					Brightness:  55,
					Temperature: 330,
				},
			},
		},
		&elgato.LightsConfig{
			NumberOfLights: 1,
			Lights: []elgato.LightState{
				{
					On:          1,
					Brightness:  35,
					Temperature: 295,
				},
			},
		},
		&elgato.LightsConfig{
			NumberOfLights: 1,
			Lights: []elgato.LightState{
				{
					On:          1,
					Brightness:  15,
					Temperature: 260,
				},
			},
		},
		&elgato.LightsConfig{
			NumberOfLights: 1,
			Lights: []elgato.LightState{
				{
					On:          0,
					Brightness:  0,
					Temperature: 225,
				},
			},
		},
	}
	for i, want := range tests {
		controller.RunEvent([]*device.Device{defaultDevice}, event)

		got := deviceServer.GetLightsConfig()
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("iteration %d failed: (-want,+got)\n%s", i, diff)
		}
	}
}
