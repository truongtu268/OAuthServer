package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/truongtu268/OAuthServer/Domain"
	"github.com/truongtu268/OAuthServer/Common"
	"github.com/truongtu268/OAuthServer/Controller"
)

func main() {
	e := echo.New()
	customValidator := Common.NewCustomValidator()
	e.Validator = customValidator
	unit := new(Domain.UnitOfWork)
	unit.Run()
	e.Use(middleware.Logger())
	mediator := new(Controller.ControllerMediator)
	mediator.InitialMediator(e)
	mediator.Execute()
	e.Logger.Fatal(e.Start(":3000"))
}
