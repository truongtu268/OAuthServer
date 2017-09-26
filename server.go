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
	Config := <- Domain.GetConfigFile()
	unit := new(Domain.UnitOfWork)
	unit.Config = Config
	unit.Run()
	e.Use(middleware.Logger())
	mediator := new(Controller.ControllerMediator)
	mediator.InitialMediator(e,Config)
	mediator.Execute()
	e.Logger.Fatal(e.Start(":9090"))
}
