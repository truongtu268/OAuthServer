package Controller

import (
	"github.com/truongtu268/OAuthServer/Middleware"
	"github.com/labstack/echo"
	"net/http"
	"github.com/truongtu268/OAuthServer/Service"
)

var service *Service.OAuthService
var oauthCotrollerItem = []ControllerItem{
	ControllerItem{
		Url:    "authorize/client",
		Method: "Post",
		HandlerFunc: func(context echo.Context) error {
			return service.OAuthClientService(context)
		},
	},
	ControllerItem{
		Url:    "authorize/user",
		Method: "Post",
		HandlerFunc: func(context echo.Context) error {
			return service.OAuthUserService(context)
		},
	},
	ControllerItem{
		Url:    "access_token",
		Method: "Get",
		HandlerFunc: func(context echo.Context) error {
			return context.JSON(http.StatusOK, "token access")
		},
	},
}

func NewOAuthController(e *echo.Echo, validatorLocate *NewMiddleware.ValidatorLocate) *EntityController {
	service = Service.NewOAuthService()
	entityCtrl := new(EntityController)
	entityCtrl.intialEntityController(e, "oauth", validatorLocate)
	for _, ctrlItem := range oauthCotrollerItem {
		entityCtrl.AddCtrlItem(ctrlItem)
	}
	return entityCtrl
}
