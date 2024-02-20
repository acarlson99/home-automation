package device_controller

type Brightness interface {
	GetBrightness() int
	SetBrightness() int
}

// ColorTemperature specifies the temperature of the color (0 = more blue increasing to red)
type ColorTemperature interface {
	GetColorTemperature() int
	SetColorTemperature() int
}
