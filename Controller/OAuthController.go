package Controller

import (
	"github.com/labstack/echo"
	"github.com/truongtu268/OAuthServer/Domain"
	"github.com/truongtu268/OAuthServer/Service"
)
var OAuthService *Service.OAuthService
var state string
var UsersControllerItem = []ControllerItem{
	ControllerItem{
		Url:"auth/github",
		Method:"Get",
		HandlerFunc: func(context echo.Context) error {
			githubProvider := OAuthService.GetService("github")
			return githubProvider.OAuthFunc(context)
		},
	},
	ControllerItem{
		Url:"login/github",
		Method:"Get",
		HandlerFunc: func(context echo.Context) error {
			githubProvider := OAuthService.GetService("github")
			return githubProvider.LoginFunc(context)
		},
	},
	ControllerItem{
		Url:"auth/google",
		Method:"Get",
		HandlerFunc: func(context echo.Context) error {
			githubProvider := OAuthService.GetService("google")
			return githubProvider.OAuthFunc(context)
		},
	},
	ControllerItem{
		Url:"login/google",
		Method:"Get",
		HandlerFunc: func(context echo.Context) error {
			githubProvider := OAuthService.GetService("google")
			return githubProvider.LoginFunc(context)
		},
	},
}
func NewOAuthController(e *echo.Echo, conf Domain.DataConfig) *EntityController {
	entityCtrl := new(EntityController)
	entityCtrl.intialEntityController(e,"user")
	OAuthService = Service.NewOAuthService(conf)
	for _, ctrlItem := range UsersControllerItem {
		entityCtrl.AddCtrlItem(ctrlItem)
	}
	return entityCtrl
}
