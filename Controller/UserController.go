package Controller

import (
	"github.com/labstack/echo"
	"github.com/truongtu268/OAuthServer/Service"
	"github.com/truongtu268/OAuthServer/Middleware"
)

var providerService *Service.ProviderService
var state string
var usersControllerItem = []ControllerItem{
	ControllerItem{
		Url:    "auth/github",
		Method: "Get",
		HandlerFunc: func(context echo.Context) error {
			githubProvider := providerService.GetService("github")
			return githubProvider.OAuthFunc(context)
		},
	},
	ControllerItem{
		Url:    "login/github",
		Method: "Get",
		HandlerFunc: func(context echo.Context) error {
			githubProvider := providerService.GetService("github")
			return githubProvider.LoginFunc(context)
		},
	},
	ControllerItem{
		Url:    "auth/google",
		Method: "Get",
		HandlerFunc: func(context echo.Context) error {
			githubProvider := providerService.GetService("google")
			return githubProvider.OAuthFunc(context)
		},
	},
	ControllerItem{
		Url:    "login/google",
		Method: "Get",
		HandlerFunc: func(context echo.Context) error {
			githubProvider := providerService.GetService("google")
			return githubProvider.LoginFunc(context)
		},
	},
	ControllerItem{
		Url:    "auth/instagram",
		Method: "Get",
		HandlerFunc: func(context echo.Context) error {
			githubProvider := providerService.GetService("instagram")
			return githubProvider.OAuthFunc(context)
		},
	},
	ControllerItem{
		Url:    "login/instagram",
		Method: "Get",
		HandlerFunc: func(context echo.Context) error {
			githubProvider := providerService.GetService("instagram")
			return githubProvider.LoginFunc(context)
		},
	},
	ControllerItem{
		Url:    "auth/internal",
		Method: "Get",
		HandlerFunc: func(context echo.Context) error {
			githubProvider := providerService.GetService("internal")
			return githubProvider.OAuthFunc(context)
		},
	},
	ControllerItem{
		Url:    "login/internal",
		Method: "Get",
		HandlerFunc: func(context echo.Context) error {
			githubProvider := providerService.GetService("internal")
			return githubProvider.LoginFunc(context)
		},
	},
}

func NewUserController(e *echo.Echo, validatorLocate *NewMiddleware.ValidatorLocate) *EntityController {
	entityCtrl := new(EntityController)
	entityCtrl.intialEntityController(e, "user", validatorLocate)
	providerService = Service.NewProviderService()

	for _, ctrlItem := range usersControllerItem {
		entityCtrl.AddCtrlItem(ctrlItem)
	}
	return entityCtrl
}
