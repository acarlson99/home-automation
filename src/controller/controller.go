package controller

import (
	"fmt"

	"github.com/acarlson99/home-automation/src/common"

	hpb "github.com/acarlson99/home-automation/proto/go"
)

func (d *Device) ExecuteAll(actions []*hpb.Event_Action) error {
	d.d.BeginBatch()
	err := common.ConcurrentAggregateErrorFn(d.Execute, actions...)
	if err != nil {
		return err
	}
	return d.d.SendBatch()
}

func (d *Device) Execute(action *hpb.Event_Action) error {
	switch act := action.Action.(type) {
	case *hpb.Event_Action_On:
		switch d := d.d.(type) {
		case PowerState:
			_, err := d.SetPowerState(act.On)
			return err
		}
	case *hpb.Event_Action_Brightness:
		switch d := d.d.(type) {
		case Brightness:
			_, err := d.SetBrightness(int(act.Brightness))
			return err
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
			_, err := d.SetColorTemperature(int(act.ColorTemp))
			return err
		}
	}
	return fmt.Errorf("action type %T called for invalid device %T", action.Action, d.GetName())
}
