package elgato

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sync"

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
	config *hpb.ElgatoLight
	client *http.Client

	mu sync.Mutex
}

func NewLight(conf *hpb.ElgatoLight) *Light {
	return &Light{
		client: http.DefaultClient,
		lights: nil,
		config: conf,
	}
}

func NewLightWithClient(conf *hpb.ElgatoLight, client *http.Client) *Light {
	return &Light{
		client: client,
		lights: nil,
		config: conf,
	}
}

func (light *Light) GetLightVals() (*LightsConfig, error) {
	url := common.FmtURL(light.config.GetUrl(), light.config.GetPort(), "/elgato/lights")
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
	url := common.FmtURL(light.config.GetUrl(), light.config.GetPort(), "/elgato/lights")
	method := "PUT"

	log.Println("setting light to cfg:", cfg)
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
	log.Println("has set light to cfg:", lightsConfig)
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
