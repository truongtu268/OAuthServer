package Controller

import (
	"github.com/labstack/echo"
	"github.com/truongtu268/OAuthServer/Domain"
)

type ControllerMediator struct {
	EntityControllers []*EntityController
	Config Domain.DataConfig
}

func (mediator *ControllerMediator)InitialMediator(echo *echo.Echo, config Domain.DataConfig) {
	mediator.EntityControllers = append(mediator.EntityControllers,
		NewOAuthController(echo, config))
}

func (mediator *ControllerMediator)Execute() error {
	for _, ctrl := range mediator.EntityControllers {
		ctrl.Execute()
	}
	return nil
}