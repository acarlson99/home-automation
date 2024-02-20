package device_controller

import (
	"fmt"

	hpb "github.com/acarlson99/home-automation/proto/go"
)

func (d *Device) ExecuteAll(actions []*hpb.Event_Action) error {
	for _, a := range actions {
		err := d.Execute(a)
		if err != nil {
			return err // TODO: aggregate errors
		}
	}
	return nil
}

func (d *Device) Execute(action *hpb.Event_Action) error {
	switch act := action.Action.(type) {
	case *hpb.Event_Action_PowerState:
		switch d := d.d.(type) {
		case PowerState:
			if act.PowerState == hpb.PowerState_Off {
				d.SetPowerState(0)
			} else {
				d.SetPowerState(1)
			}
		}
	case *hpb.Event_Action_Brightness:
		switch d := d.d.(type) {
		case Brightness:
			d.SetBrightness(int(act.Brightness))
		}
	case *hpb.Event_Action_RelativeBrightness:
		return fmt.Errorf("relative brightness not implemented")
	case *hpb.Event_Action_Color:
		return fmt.Errorf("color not implemented")
		// switch d := d.d.(type) {
		// case Color:
		// 	d.SetColor(int(act.Color))
		// }
	case *hpb.Event_Action_ColorTemp:
		switch d := d.d.(type) {
		case ColorTemperature:
			d.SetColorTemperature(int(act.ColorTemp))
		}
	default:
		return fmt.Errorf("action type %T called for invalid device %T", act, d)
	}
	return nil
}
