package testutil

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/acarlson99/home-automation/src/common"
	"github.com/acarlson99/home-automation/src/device"
	"github.com/acarlson99/home-automation/src/elgato"
)

type FakeLightServer struct {
	*httptest.Server
	lightsConfig *elgato.LightsConfig
	name         string
}

var (
	_a = device.IDevice(&FakeLightServer{})
	_b = device.PowerState(&FakeLightServer{})
	_c = device.ColorTemperature(&FakeLightServer{})
	_d = device.Brightness(&FakeLightServer{})
)

// GetBrightness implements device.Brightness.
func (fs *FakeLightServer) GetBrightness() (int, error) {
	return fs.lightsConfig.Lights[0].Brightness, nil
}

// GetMaxBrightness implements device.Brightness.
func (*FakeLightServer) GetMaxBrightness() int {
	return elgato.MaxBrightness
}

// GetMinBrightness implements device.Brightness.
func (*FakeLightServer) GetMinBrightness() int {
	return elgato.MinBrightness
}

// SetBrightness implements device.Brightness.
func (fs *FakeLightServer) SetBrightness(n int) (int, error) {
	fs.lightsConfig.Lights[0].Brightness = n
	return fs.lightsConfig.Lights[0].Brightness, nil
}

// GetColorTemperature implements device.ColorTemperature.
func (fs *FakeLightServer) GetColorTemperature() (int, error) {
	return fs.lightsConfig.Lights[0].Temperature, nil
}

// GetMaxColorTemperature implements device.ColorTemperature.
func (*FakeLightServer) GetMaxColorTemperature() int {
	return elgato.MaxColorTemperature
}

// GetMinColorTemperature implements device.ColorTemperature.
func (*FakeLightServer) GetMinColorTemperature() int {
	return elgato.MinColorTemperature
}

// SetColorTemperature implements device.ColorTemperature.
func (fs *FakeLightServer) SetColorTemperature(n int) (int, error) {
	fs.lightsConfig.Lights[0].Temperature = n
	return fs.lightsConfig.Lights[0].Temperature, nil
}

// GetPowerState implements device.PowerState.
func (fs *FakeLightServer) GetPowerState() (bool, error) {
	return fs.lightsConfig.Lights[0].On == 1, nil
}

// SetPowerState implements device.PowerState.
func (fs *FakeLightServer) SetPowerState(b bool) (bool, error) {
	if b {
		fs.lightsConfig.Lights[0].On = 1
	} else {
		fs.lightsConfig.Lights[0].On = 0
	}
	return fs.lightsConfig.Lights[0].On == 1, nil
}

func (fs *FakeLightServer) GetLightsConfig() *elgato.LightsConfig {
	return fs.lightsConfig
}

func (fs *FakeLightServer) SetLightsConfig(cfg elgato.LightsConfig) *elgato.LightsConfig {
	fs.lightsConfig = &cfg
	return fs.lightsConfig
}

// BeginBatch implements device.IDevice.
func (*FakeLightServer) BeginBatch() error {
	return nil
}

// GetName implements device.IDevice.
func (s *FakeLightServer) GetName() string {
	return s.name
}

// NameMatches implements device.IDevice.
func (fs *FakeLightServer) NameMatches(s string) bool {
	return s == fs.GetName()
}

// SendBatch implements device.IDevice.
func (*FakeLightServer) SendBatch() error {
	return nil
}

func NewFakeLightServer(t *testing.T, name string, conf *elgato.LightsConfig) (*FakeLightServer, func()) {
	ls := &FakeLightServer{lightsConfig: conf, name: name}
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
			reqLights := &elgato.LightsConfig{}
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

func (s *FakeLightServer) Handle(reqLights *elgato.LightsConfig) {
	// assume length of 1
	if len(reqLights.Lights) < 1 {
		return
	}
	l := reqLights.Lights[0]
	if want, v := l.On, common.Clamp(0, l.On, 1); v == want {
		s.lightsConfig.Lights[0].On = v
	}
	if want, v := l.Brightness, common.Clamp(elgato.MinBrightness, l.Brightness, elgato.MaxBrightness); v == want {
		s.lightsConfig.Lights[0].Brightness = v
	}
	if want, v := l.Temperature, common.Clamp(elgato.MinColorTemperature, l.Temperature, elgato.MaxColorTemperature); v == want {
		s.lightsConfig.Lights[0].Temperature = v
	}
}
