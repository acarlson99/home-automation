package elgato

import (
	"github.com/acarlson99/home-automation/src/device"
)

var (
	_a = device.IDevice(&Light{})
	_b = device.PowerState(&Light{})
	_c = device.ColorTemperature(&Light{})
	_d = device.Brightness(&Light{})
)

// implement controller.IDevice
func (light *Light) GetName() string {
	return light.config.Name
}

func (light *Light) NameMatches(s string) bool {
	return light.GetName() == s
}

func (light *Light) EndBatch() {
	light.mu.Unlock()
}

func (light *Light) BeginBatch() error {
	light.mu.Lock()
	vs, err := light.GetLightVals()
	if err != nil {
		return err
	}
	light.lights = vs
	return nil
}

func (light *Light) SendBatch() error {
	_, err := light.SetLightVals(light.lights)
	if err != nil {
		return err
	}
	return nil
}

func (light *Light) GetMaxBrightness() int {
	return MaxBrightness
}

func (light *Light) GetMinBrightness() int {
	return MinBrightness
}

func (light *Light) GetBrightness() (int, error) {
	// vs, err := light.GetLightVals()
	// if err != nil {
	// 	return 0, err
	// }
	// if len(vs.Lights) < 1 {
	// 	return 0, fmt.Errorf("light.GetLightVals() unexpectedly returned `%+v`", vs)
	// }
	return light.lights.Lights[0].Brightness, nil
}

func (light *Light) SetBrightness(n int) (int, error) {
	// light.mu.Lock()
	// defer light.mu.Unlock()
	// common.Logger(common.Debug).Println("set brightness to", n)
	// vs, err := light.GetLightVals()
	// if err != nil {
	// 	return 0, err
	// }
	// if len(vs.Lights) < 1 {
	// 	return 0, fmt.Errorf("light.GetLightVals() unexpectedly returned `%+v`", vs)
	// }
	for i := range light.lights.Lights {
		light.lights.Lights[i].Brightness = n
	}

	// vs, err = light.SetLightVals(vs)
	// if err != nil {
	// 	return 0, err
	// }
	// if len(vs.Lights) < 1 {
	// 	return 0, fmt.Errorf("light.GetLightVals() unexpectedly returned `%+v`", vs)
	// }

	return light.GetBrightness()
}

func (light *Light) GetMaxColorTemperature() int {
	return MaxColorTemperature
}

func (light *Light) GetMinColorTemperature() int {
	return MinColorTemperature
}

func (light *Light) GetColorTemperature() (int, error) {
	// vs, err := light.GetLightVals()
	// if err != nil {
	// 	return 0, err
	// }
	// if len(vs.Lights) < 1 {
	// 	return 0, fmt.Errorf("light.GetLightVals() unexpectedly returned `%+v`", vs)
	// }
	return light.lights.Lights[0].Temperature, nil
}

func (light *Light) SetColorTemperature(n int) (int, error) {
	// light.mu.Lock()
	// defer light.mu.Unlock()
	// vs, err := light.GetLightVals()
	// if err != nil {
	// 	return 0, err
	// }
	// if len(vs.Lights) < 1 {
	// 	return 0, fmt.Errorf("light.GetLightVals() unexpectedly returned `%+v`", vs)
	// }
	for i := range light.lights.Lights {
		light.lights.Lights[i].Temperature = n
	}

	// vs, err = light.SetLightVals(vs)
	// if err != nil {
	// 	return 0, err
	// }
	// if len(vs.Lights) < 1 {
	// 	return 0, fmt.Errorf("light.GetLightVals() unexpectedly returned `%+v`", vs)
	// }

	return light.GetColorTemperature()
}

func (light *Light) GetPowerState() (bool, error) {
	// vs, err := light.GetLightVals()
	// if err != nil {
	// 	return false, err
	// }
	// if len(vs.Lights) < 1 {
	// 	return false, fmt.Errorf("light.GetLightVals() unexpectedly returned `%+v`", vs)
	// }
	return light.lights.Lights[0].On == 1, nil
}

func (light *Light) SetPowerState(on bool) (bool, error) {
	// vs, err := light.GetLightVals()
	// if err != nil {
	// 	return false, err
	// }
	// if len(vs.Lights) < 1 {
	// 	return false, fmt.Errorf("light.GetLightVals() unexpectedly returned `%+v`", vs)
	// }
	for i := range light.lights.Lights {
		if on {
			light.lights.Lights[i].On = 1
		} else {
			light.lights.Lights[i].On = 0
		}
	}

	// vs, err = light.SetLightVals(vs)
	// if err != nil {
	// 	return false, err
	// }
	// if len(vs.Lights) < 1 {
	// 	return false, fmt.Errorf("light.GetLightVals() unexpectedly returned `%+v`", vs)
	// }

	return light.GetPowerState()
}
