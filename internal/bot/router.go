package bot

import "github.com/AlexeyNilov/zero_control/internal/service"

type Router struct {
	control *service.ControlService
}

func NewRouter(control *service.ControlService) Router {
	return Router{control: control}
}

func (r Router) StatusMessage() string {
	return "zero_control is online"
}
