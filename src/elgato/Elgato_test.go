package elgato

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"

	hpb "github.com/acarlson99/home-automation/proto/go"
)

func clamp(a, b, c int) int {
	return min(c, max(a, b))
	// return min(max(a, b), c)
}

type FakeLightServer struct {
	*httptest.Server
	lightsConfig *LightsConfig
}

func NewFakeLightServer(t *testing.T, conf *LightsConfig) (*FakeLightServer, func()) {
	ls := &FakeLightServer{lightsConfig: conf}
	server := httptest.NewServer(http.HandlerFunc(ls.HandlerFunc(t)))
	ls.Server = server
	return ls, ls.Close
}

func (s *FakeLightServer) HandlerFunc(t *testing.T) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		if req.Method == "PUT" {
			body, err := io.ReadAll(req.Body)
			if err != nil {
				t.Errorf("Invalid request `%s`: err %v", string(body), err)
			}
			reqLights := &LightsConfig{}
			json.Unmarshal(body, reqLights)
			// log.Println(reqLights)

			s.Handle(reqLights)
		}
		r, err := json.Marshal(s.lightsConfig)
		if err != nil {
			t.Fatal(err)
		}
		rw.Write(r)
	}
}

func (s *FakeLightServer) Handle(reqLights *LightsConfig) {
	// assume length of 1
	if len(reqLights.Lights) < 1 {
		return
	}
	l := reqLights.Lights[0]
	if want, v := l.On, clamp(0, l.On, 1); v == want {
		s.lightsConfig.Lights[0].On = v
	}
	if want, v := l.Brightness, clamp(MinBrightness, l.Brightness, MaxBrightness); v == want {
		s.lightsConfig.Lights[0].Brightness = v
	}
	if want, v := l.Temperature, clamp(MinTemp, l.Temperature, MaxTemp); v == want {
		s.lightsConfig.Lights[0].Temperature = v
	}
}

func TestSetValsTestWithServer(t *testing.T) {
	defaultTemp := 200
	conf := LightsConfig{
		NumberOfLights: 1,
		Lights: []LightState{
			{
				On:          0,
				Brightness:  50,
				Temperature: defaultTemp,
			},
		},
	}
	server, cleanup := NewFakeLightServer(t, &conf)
	client := server.Client()
	url := server.URL
	defer cleanup()

	for _, tt := range []struct {
		name    string
		url     string
		arg     *LightsConfig
		wantErr bool
		want    *LightsConfig
	}{
		{
			name: "below min",
			url:  url,
			arg: &LightsConfig{
				NumberOfLights: 1,
				Lights: []LightState{
					{
						On:          1,
						Brightness:  20,
						Temperature: 100,
					},
				},
			},
			want: &LightsConfig{NumberOfLights: 1,
				Lights: []LightState{
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
			arg: &LightsConfig{
				NumberOfLights: 1,
				Lights: []LightState{
					{
						On:          0,
						Brightness:  99,
						Temperature: 300,
					},
				},
			},
			want: &LightsConfig{NumberOfLights: 1,
				Lights: []LightState{
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
			arg: &LightsConfig{
				NumberOfLights: 1,
				Lights: []LightState{
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
			light := NewLightWithClient(lightpb, client)
			light.SetLightVals(&LightsConfig{NumberOfLights: 1, Lights: []LightState{{On: 0, Brightness: 50, Temperature: defaultTemp}}})
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
	conf := LightsConfig{
		NumberOfLights: 1,
		Lights: []LightState{
			{
				On:          0,
				Brightness:  50,
				Temperature: defaultTemp,
			},
		},
	}
	server, cleanup := NewFakeLightServer(t, &conf)
	client := server.Client()
	url := server.URL
	defer cleanup()

	for _, tt := range []struct {
		name    string
		url     string
		init    *LightsConfig
		arg     *LightsConfig
		wantErr bool
		want    *LightsConfig
	}{
		{
			name: "adding above max does not change",
			url:  url,
			init: &LightsConfig{
				NumberOfLights: 1,
				Lights: []LightState{
					{
						On:          1,
						Brightness:  50,
						Temperature: 200,
					},
				},
			},
			arg: &LightsConfig{
				NumberOfLights: 1,
				Lights: []LightState{
					{
						On:          1,
						Brightness:  50,
						Temperature: 100,
					},
				},
			},
			want: &LightsConfig{NumberOfLights: 1,
				Lights: []LightState{
					{
						On:          1,
						Brightness:  MaxBrightness,
						Temperature: 300,
					},
				},
			},
		},
		{
			name: "subtract",
			url:  url,
			init: &LightsConfig{
				NumberOfLights: 1,
				Lights: []LightState{
					{
						On:          1,
						Brightness:  99,
						Temperature: 300,
					},
				},
			},
			arg: &LightsConfig{
				NumberOfLights: 1,
				Lights: []LightState{
					{
						On:          1,
						Brightness:  -96,
						Temperature: -150,
					},
				},
			},
			want: &LightsConfig{NumberOfLights: 1,
				Lights: []LightState{
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
			init: &LightsConfig{
				NumberOfLights: 1,
				Lights: []LightState{
					{
						On:          1,
						Brightness:  99,
						Temperature: 300,
					},
				},
			},
			arg: &LightsConfig{
				NumberOfLights: 1,
				Lights: []LightState{
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
			light := NewLightWithClient(lightpb, client)
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
