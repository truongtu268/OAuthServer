package Controller

import (
	"github.com/labstack/echo"
)

type ControllerMediator struct {
	EntityControllers []*EntityController
}

func (mediator *ControllerMediator)InitialMediator(echo *echo.Echo) {
	mediator.EntityControllers = append(mediator.EntityControllers,
		NewOAuthController(echo))
}

func (mediator *ControllerMediator)Execute() error {
	for _, ctrl := range mediator.EntityControllers {
		ctrl.Execute()
	}
	return nil
}