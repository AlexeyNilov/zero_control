package device

import "context"

type Device struct{}

func New() *Device {
	return &Device{}
}

func (d *Device) Status(context.Context) (string, error) {
	return "unknown", nil
}
