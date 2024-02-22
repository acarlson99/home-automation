package elgato_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	hpb "github.com/acarlson99/home-automation/proto/go"
	"github.com/acarlson99/home-automation/src/elgato"
	"github.com/acarlson99/home-automation/src/elgato/testutil"
)

func TestSetValsTestWithServer(t *testing.T) {
	defaultTemp := 200
	conf := elgato.LightsConfig{
		NumberOfLights: 1,
		Lights: []elgato.LightState{
			{
				On:          0,
				Brightness:  50,
				Temperature: defaultTemp,
			},
		},
	}
	server, cleanup := testutil.NewFakeLightServer(t, "test-light", &conf)
	client := server.Client()
	url := server.URL
	defer cleanup()

	for _, tt := range []struct {
		name    string
		url     string
		arg     *elgato.LightsConfig
		wantErr bool
		want    *elgato.LightsConfig
	}{
		{
			name: "below min",
			url:  url,
			arg: &elgato.LightsConfig{
				NumberOfLights: 1,
				Lights: []elgato.LightState{
					{
						On:          1,
						Brightness:  20,
						Temperature: 100,
					},
				},
			},
			want: &elgato.LightsConfig{NumberOfLights: 1,
				Lights: []elgato.LightState{
					{
						On:          1,
						Brightness:  20,
						Temperature: defaultTemp,
					},
				},
			},
		},
		{
			name: "good set",
			url:  url,
			arg: &elgato.LightsConfig{
				NumberOfLights: 1,
				Lights: []elgato.LightState{
					{
						On:          0,
						Brightness:  99,
						Temperature: 300,
					},
				},
			},
			want: &elgato.LightsConfig{NumberOfLights: 1,
				Lights: []elgato.LightState{
					{
						On:          0,
						Brightness:  99,
						Temperature: 300,
					},
				},
			},
		},
		{
			name: "bad url should error gracefully",
			url:  "hhp:t//cumdumpst.er",
			arg: &elgato.LightsConfig{
				NumberOfLights: 1,
				Lights: []elgato.LightState{
					{
						On:          0,
						Brightness:  99,
						Temperature: 300,
					},
				},
			},
			wantErr: true,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			lightpb := &hpb.ElgatoLight{Url: tt.url}
			light := elgato.NewLightWithClient(lightpb, client)
			light.SetLightVals(&elgato.LightsConfig{NumberOfLights: 1, Lights: []elgato.LightState{{On: 0, Brightness: 50, Temperature: defaultTemp}}})

			got, err := light.SetLightVals(tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetLightVals() error = %v, wantErr %v", err, tt.wantErr)
			} else if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("set return mismatch (-want, +got)\n%s", diff)
			}

			got, err = light.GetLightVals()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLightVals() error = %v, wantErr %v", err, tt.wantErr)
			} else if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("get return mismatch (-want, +got)\n%s", diff)
			}
		})
	}
}

func TestChangeValsTestWithServer(t *testing.T) {
	defaultTemp := 200
	conf := elgato.LightsConfig{
		NumberOfLights: 1,
		Lights: []elgato.LightState{
			{
				On:          0,
				Brightness:  50,
				Temperature: defaultTemp,
			},
		},
	}
	server, cleanup := testutil.NewFakeLightServer(t, "test-light", &conf)
	client := server.Client()
	url := server.URL
	defer cleanup()

	for _, tt := range []struct {
		name    string
		url     string
		init    *elgato.LightsConfig
		arg     *elgato.LightsConfig
		wantErr bool
		want    *elgato.LightsConfig
	}{
		{
			name: "adding above max does not change",
			url:  url,
			init: &elgato.LightsConfig{
				NumberOfLights: 1,
				Lights: []elgato.LightState{
					{
						On:          1,
						Brightness:  50,
						Temperature: 200,
					},
				},
			},
			arg: &elgato.LightsConfig{
				NumberOfLights: 1,
				Lights: []elgato.LightState{
					{
						On:          1,
						Brightness:  50,
						Temperature: 100,
					},
				},
			},
			want: &elgato.LightsConfig{NumberOfLights: 1,
				Lights: []elgato.LightState{
					{
						On:          1,
						Brightness:  elgato.MaxBrightness,
						Temperature: 300,
					},
				},
			},
		},
		{
			name: "subtract",
			url:  url,
			init: &elgato.LightsConfig{
				NumberOfLights: 1,
				Lights: []elgato.LightState{
					{
						On:          1,
						Brightness:  99,
						Temperature: 300,
					},
				},
			},
			arg: &elgato.LightsConfig{
				NumberOfLights: 1,
				Lights: []elgato.LightState{
					{
						On:          1,
						Brightness:  -96,
						Temperature: -150,
					},
				},
			},
			want: &elgato.LightsConfig{NumberOfLights: 1,
				Lights: []elgato.LightState{
					{
						On:          1,
						Brightness:  3,
						Temperature: 150,
					},
				},
			},
		},
		{
			name: "empty url for change should error gracefully",
			url:  "",
			init: &elgato.LightsConfig{
				NumberOfLights: 1,
				Lights: []elgato.LightState{
					{
						On:          1,
						Brightness:  99,
						Temperature: 300,
					},
				},
			},
			arg: &elgato.LightsConfig{
				NumberOfLights: 1,
				Lights: []elgato.LightState{
					{
						On:          1,
						Brightness:  -96,
						Temperature: -150,
					},
				},
			},
			wantErr: true,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			lightpb := &hpb.ElgatoLight{Url: tt.url}
			light := elgato.NewLightWithClient(lightpb, client)
			light.SetLightVals(tt.init)

			got, err := light.ChangeLightVals(tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChangeLightVals() error = %v, wantErr %v", err, tt.wantErr)
			} else if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("set return mismatch (-want, +got)\n%s", diff)
			}

			got, err = light.GetLightVals()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLightVals() error = %v, wantErr %v", err, tt.wantErr)
			} else if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("get return mismatch (-want, +got)\n%s", diff)
			}
		})
	}
}
