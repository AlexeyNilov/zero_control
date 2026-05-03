package service

import "context"

type Device interface {
	Status(context.Context) (string, error)
}

type ControlService struct {
	device Device
}

func New(device Device) *ControlService {
	return &ControlService{device: device}
}

func (s *ControlService) Status(ctx context.Context) (string, error) {
	return s.device.Status(ctx)
}
