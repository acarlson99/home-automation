package device

import (
	"fmt"
	"log"
	"sync"

	hpb "github.com/acarlson99/home-automation/proto/go"
	"github.com/acarlson99/home-automation/src/common"
	"golang.org/x/exp/constraints"
)

var (
	ds sync.Map
)

type IDevice interface {
	GetName() string
	NameMatches(s string) bool

	BeginBatch() error // batch a set of actions
	SendBatch() error  // send all actions made since `BeginBatch`
}

type Device struct {
	d IDevice
}

func NewDevice(d IDevice) *Device {
	return &Device{d: d}
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
	GetMaxBrightness() int
	GetMinBrightness() int
}

// ColorTemperature specifies the temperature of the color (0 = more blue increasing to red)
type ColorTemperature interface {
	GetColorTemperature() (int, error)
	SetColorTemperature(int) (int, error)
	GetMaxColorTemperature() int
	GetMinColorTemperature() int
}

type PowerState interface {
	GetPowerState() (bool, error)
	SetPowerState(bool) (bool, error)
}

func GetDevices() []*Device {
	l := []*Device{}
	ds.Range(func(k, v any) bool {
		l = append(l, v.(*Device))
		return true
	})
	return l
}

func RegisterDevice(d *Device) error {
	ds.Store(d.GetName(), d)
	return nil
}

func UnregisterDevice(d *Device) error {
	ds.Delete(d.GetName())
	return nil
}

func (d *Device) IdentifierMatches(ident *hpb.DeviceIdentifier) bool {
	return d.NameMatches(ident.GetName())
}

func (d *Device) GetDeviceState(t *hpb.DeviceState_Type) (*hpb.Primitive, error) {
	switch *t {
	case hpb.DeviceState_Power:
		pd, ok := d.d.(PowerState)
		if !ok {
			return nil, fmt.Errorf("state request %v invalid for device type %T", *t, d)
		}
		b, err := pd.GetPowerState()
		if err != nil {
			return nil, fmt.Errorf("unable to get state %v for device type %T", *t, err)
		}
		return &hpb.Primitive{V: &hpb.Primitive_Bool{Bool: b}}, nil
	case hpb.DeviceState_Brightness:
		pd, ok := d.d.(Brightness)
		if !ok {
			return nil, fmt.Errorf("state request %v invalid for device type %T", *t, d)
		}
		n, err := pd.GetBrightness()
		if err != nil {
			return nil, fmt.Errorf("unable to get state %v for device type %T", *t, err)
		}
		return &hpb.Primitive{V: &hpb.Primitive_Int32{Int32: int32(n)}}, nil
	// case hpb.DeviceState_Color:
	// 	return nil, fmt.Errorf("state request %v not implemented", *t)
	case hpb.DeviceState_ColorTemperature:
		pd, ok := d.d.(ColorTemperature)
		if !ok {
			return nil, fmt.Errorf("state request %v invalid for device type %T", *t, d)
		}
		n, err := pd.GetColorTemperature()
		if err != nil {
			return nil, fmt.Errorf("unable to get state %v for device type %T", *t, err)
		}
		return &hpb.Primitive{V: &hpb.Primitive_Int32{Int32: int32(n)}}, nil
	}
	return nil, fmt.Errorf("state request %v matched nothing for device type %T", *t, d)
}

func (d *Device) ExecuteAll(actions []*hpb.Event_Action) error {
	d.d.BeginBatch()
	err := common.AggregateErrorFn(d.Execute, actions...)
	if err != nil {
		return err
	}
	return d.d.SendBatch()
}

// percent should be number from 0-1
func percentStep[T constraints.Float | constraints.Integer](percent float32, max, min T) T {
	return T(percent * float32(max-min))
}

func (dev *Device) Execute(action *hpb.Event_Action) error {
	switch act := action.Action.(type) {
	case *hpb.Event_Action_On:
		switch d := dev.d.(type) {
		case PowerState:
			_, err := d.SetPowerState(act.On)
			return err
		}
	case *hpb.Event_Action_Brightness:
		switch d := dev.d.(type) {
		case Brightness:
			brightness := act.Brightness
			if action.GetRelative() {
				current, err := d.GetBrightness()
				if err != nil {
					return err
				}
				log.Println("current brightness", current, "changing", act.Brightness, "%")
				brightness = int32(current + percentStep(float32(act.Brightness)/100.0, d.GetMaxBrightness(), d.GetMinBrightness()))
				log.Println("modifying brightness to", act.Brightness)
			}
			_, err := d.SetBrightness(common.Clamp(d.GetMinBrightness(), int(brightness), d.GetMaxBrightness()))
			if err != nil {
				return err
			}
			if brightness <= 0 {
				powerDevice, ok := dev.d.(PowerState)
				if !ok {
					return nil
				}
				_, err = powerDevice.SetPowerState(false)
				return err
			}
		}
	// case *hpb.Event_Action_Color:
	// 	return fmt.Errorf("color not implemented")
	// switch d := d.d.(type) {
	// case Color:
	// 	d.SetColor(int(act.Color))
	// }
	case *hpb.Event_Action_ColorTemp:
		switch d := dev.d.(type) {
		case ColorTemperature:
			colorTemp := act.ColorTemp
			if action.GetRelative() {
				current, err := d.GetColorTemperature()
				if err != nil {
					return err
				}
				colorTemp = int32(current + percentStep(float32(act.ColorTemp)/100.0, d.GetMaxColorTemperature(), d.GetMinColorTemperature()))
			}
			_, err := d.SetColorTemperature(common.Clamp(d.GetMinColorTemperature(), int(colorTemp), d.GetMaxColorTemperature()))
			return err
		}
	}
	return fmt.Errorf("action type %T called for invalid device %s", action.Action, dev.GetName())
}
