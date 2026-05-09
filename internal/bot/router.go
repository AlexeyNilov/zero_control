package bot

import (
	"context"
	"strings"

	"github.com/AlexeyNilov/zero_control/internal/service"
)

type Router struct {
	control *service.ControlService
}

func NewRouter(control *service.ControlService) Router {
	return Router{control: control}
}

func (r Router) StatusMessage(ctx context.Context) string {
	ipAddress := "unavailable"
	if r.control != nil {
		status, err := r.control.Status(ctx)
		if err == nil && strings.TrimSpace(status) != "" {
			ipAddress = status
		}
	}

	return "zero_control is online\nip: " + ipAddress
}
