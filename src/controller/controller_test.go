package controller_test

import (
	"log"
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
	lname := "test-light"
	lightServer, cleanup := testutil.NewFakeLightServer(t, lname, nil)
	defaultDevice := device.NewDevice(lightServer)
	defer cleanup()
	device.RegisterDevice(defaultDevice)
	defer device.UnregisterDevice(defaultDevice)

	defaultLightConf := elgato.LightState{
		On:          1,
		Brightness:  50,
		Temperature: 50,
	}

	type args struct {
		event string
	}
	tests := []struct {
		name string
		args args
		want *elgato.LightsConfig
	}{
		{
			name: "static set brightness",
			args: args{
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
						Temperature: 50 - 20,
					},
				},
			},
		},
		{
			name: "change brightness successful conditional",
			args: args{
				event: `name: "conditional"
devices {
	name: "test-light"
}
actions {
	brightness: 100
}
actions {
	color_temp: 100
}
start_if {
	op: EQ
	lhs {
		prim {
			bool: true
		}
	}
	rhs {
		device_state {
			device {
				name: "test-light"
			}
			type: Power
		}
	}
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
						Brightness:  100,
						Temperature: 100,
					},
				},
			},
		},
		{
			name: "change brightness with invalid conditional should yield no change",
			args: args{
				event: `name: "conditional"
devices {
	name: "test-light"
}
actions {
	brightness: 30
}
actions {
	color_temp: 20
}
start_if {
	op: EQ
	lhs {
		prim {
			int32: 30
		}
	}
	rhs {
		device_state {
			device {
				name: "test-light"
			}
			type: Power
		}
	}
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
				Lights:         []elgato.LightState{defaultLightConf},
			},
		},
		{
			name: "change brightness with failing conditional should yield no change",
			args: args{
				event: `name: "conditional"
devices {
	name: "test-light"
}
actions {
	brightness: 30
}
actions {
	color_temp: 20
}
start_if {
	op: NEQ
	lhs {
		prim {
			bool: true
		}
	}
	rhs {
		device_state {
			device {
				name: "test-light"
			}
			type: Power
		}
	}
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
				Lights:         []elgato.LightState{defaultLightConf},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lightServer.SetLightsConfig(elgato.LightsConfig{
				NumberOfLights: 1,
				Lights:         []elgato.LightState{defaultLightConf},
			})

			event := &hpb.Event{}
			err := tpb.Unmarshal([]byte(tt.args.event), event)
			if err != nil {
				t.Fatal("could not unmarshal test:", err)
			}
			log.Printf("run event %s %+v", event.Name, event.GetStartIf())
			controller.RunEvent([]*device.Device{defaultDevice}, event)

			if diff := cmp.Diff(tt.want, lightServer.GetLightsConfig()); diff != "" {
				t.Errorf("(-want, +got)\n%s", diff)
			}
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
		{
			NumberOfLights: 1,
			Lights: []elgato.LightState{
				{
					On:          1,
					Brightness:  75,
					Temperature: 365,
				},
			},
		},
		{
			NumberOfLights: 1,
			Lights: []elgato.LightState{
				{
					On:          1,
					Brightness:  55,
					Temperature: 330,
				},
			},
		},
		{
			NumberOfLights: 1,
			Lights: []elgato.LightState{
				{
					On:          1,
					Brightness:  35,
					Temperature: 295,
				},
			},
		},
		{
			NumberOfLights: 1,
			Lights: []elgato.LightState{
				{
					On:          1,
					Brightness:  15,
					Temperature: 260,
				},
			},
		},
		{
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
