package elgato

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/acarlson99/home-automation/src/common"

	hpb "github.com/acarlson99/home-automation/proto/go"
)

const (
	MinTemp       = 143
	MaxTemp       = 344
	MinBrightness = 3
	MaxBrightness = 100
)

type LightState struct {
	On          int `json:"on"`
	Brightness  int `json:"brightness"`
	Temperature int `json:"temperature"`
}

type LightsConfig struct {
	NumberOfLights int          `json:"numberOfLights"`
	Lights         []LightState `json:"lights"`
}

type Light struct {
	lights *LightsConfig
	pb     *hpb.ElgatoLight
	client *http.Client
}

func NewLight(pb *hpb.ElgatoLight) *Light {
	return &Light{
		client: http.DefaultClient,
		lights: nil,
		pb:     pb,
	}
}

func NewLightWithClient(pb *hpb.ElgatoLight, client *http.Client) *Light {
	return &Light{
		client: client,
		lights: nil,
		pb:     pb,
	}
}

func (light *Light) GetLightVals() (*LightsConfig, error) {
	url := common.FmtURL(light.pb.GetUrl(), light.pb.GetPort(), "/elgato/lights")
	method := "GET"
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	req.Header.Add("Content-Type", "text/json")

	res, err := light.client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var lightsConfig LightsConfig
	err = json.Unmarshal(body, &lightsConfig)
	return &lightsConfig, err
}

func (light *Light) SetLightVals(cfg *LightsConfig) (*LightsConfig, error) {
	url := common.FmtURL(light.pb.GetUrl(), light.pb.GetPort(), "/elgato/lights")
	method := "PUT"

	bs, err := json.Marshal(cfg)
	if err != nil {
		return nil, err
	}
	payload := bytes.NewReader(bs)

	req, err := http.NewRequest(method, url, payload)
	// req.Close = true

	if err != nil {
		log.Println(err)
		return nil, err
	}
	req.Header.Add("Content-Type", "text/json")

	res, err := light.client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var lightsConfig LightsConfig
	err = json.Unmarshal(body, &lightsConfig)
	return &lightsConfig, err
}

func (light *Light) ChangeLightVals(diffs *LightsConfig) (*LightsConfig, error) {
	states, err := light.GetLightVals()
	if err != nil {
		return nil, err
	}
	for i, c := range states.Lights {
		l := c
		l.On = diffs.Lights[i].On
		l.Brightness += diffs.Lights[i].Brightness
		l.Temperature += diffs.Lights[i].Temperature
		states.Lights[i] = l
	}
	return light.SetLightVals(states)
}
