package Controller

import (
	"github.com/labstack/echo"
	"github.com/truongtu268/OAuthServer/Middleware"
)

type ControllerMediator struct {
	EntityControllers []*EntityController
}

func (mediator *ControllerMediator)InitialMediator(echo *echo.Echo) {
	validatorLocate:= NewMiddleware.NewValidatorLocation()
	mediator.EntityControllers = append(mediator.EntityControllers,
		NewOAuthController(echo, validatorLocate),
		NewPrivateController(echo, validatorLocate),)
}

func (mediator *ControllerMediator)Execute() error {
	for _, ctrl := range mediator.EntityControllers {
		ctrl.Execute()
	}
	return nil
}