package elgato

import (
	"fmt"

	device_controller "github.com/acarlson99/home-automation/src/device-controller"
)

func NewDevice(light *Light) *device_controller.Device {
	return device_controller.NewDevice(light)
}

func (light *Light) GetBrightness() (int, error) {
	vs, err := light.GetLightVals()
	if err != nil {
		return 0, err
	}
	if len(vs.Lights) < 1 {
		return 0, fmt.Errorf("light.GetLightVals() unexpectedly returned `%+v`", vs)
	}
	return vs.Lights[0].Brightness, nil
}

func (light *Light) SetBrightness(n int) (int, error) {
	vs, err := light.GetLightVals()
	if err != nil {
		return 0, err
	}
	if len(vs.Lights) < 1 {
		return 0, fmt.Errorf("light.GetLightVals() unexpectedly returned `%+v`", vs)
	}
	for i := range vs.Lights {
		vs.Lights[i].Brightness = n
	}

	vs, err = light.SetLightVals(vs)
	if err != nil {
		return 0, err
	}
	if len(vs.Lights) < 1 {
		return 0, fmt.Errorf("light.GetLightVals() unexpectedly returned `%+v`", vs)
	}

	return light.GetBrightness()
}

func (light *Light) GetColorTemperature() (int, error) {
	vs, err := light.GetLightVals()
	if err != nil {
		return 0, err
	}
	if len(vs.Lights) < 1 {
		return 0, fmt.Errorf("light.GetLightVals() unexpectedly returned `%+v`", vs)
	}
	return vs.Lights[0].Temperature, nil
}

func (light *Light) SetColorTemperature(n int) (int, error) {
	vs, err := light.GetLightVals()
	if err != nil {
		return 0, err
	}
	if len(vs.Lights) < 1 {
		return 0, fmt.Errorf("light.GetLightVals() unexpectedly returned `%+v`", vs)
	}
	for i := range vs.Lights {
		vs.Lights[i].Temperature = n
	}

	vs, err = light.SetLightVals(vs)
	if err != nil {
		return 0, err
	}
	if len(vs.Lights) < 1 {
		return 0, fmt.Errorf("light.GetLightVals() unexpectedly returned `%+v`", vs)
	}

	return light.GetColorTemperature()
}
