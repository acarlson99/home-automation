package controller

type IDevice interface {
	GetName() string
	NameMatches(s string) bool

	BeginBatch() error // batch a set of actions
	SendBatch() error  // send all actions made since `BeginBatch`
}

func NewDevice(d IDevice) *Device {
	return &Device{d: d}
}

type Device struct {
	d IDevice
}

func (d *Device) GetName() string {
	return d.d.GetName()
}

func (d *Device) NameMatches(s string) bool {
	return d.d.NameMatches(s)
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
	GetPowerState() (bool, error)
	SetPowerState(bool) (bool, error)
}
