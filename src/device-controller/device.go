package device_controller

type Device struct {
	d interface{}
}

func NewDevice(d interface{}) *Device {
	return &Device{d: d}
}

type Brightness interface {
	GetBrightness() (int, error)
	SetBrightness(int) (int, error)
}

// ColorTemperature specifies the temperature of the color (0 = more blue increasing to red)
type ColorTemperature interface {
	GetColorTemperature() (int, error)
	SetColorTemperature(int) (int, error)
}

type PowerState interface {
	GetPowerState() (int, error)
	SetPowerState(int) (int, error)
}
