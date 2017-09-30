package Controller

import (
	"github.com/truongtu268/OAuthServer/Middleware"
	"github.com/labstack/echo"
	"net/http"
)

var oauthCotrollerItem = []ControllerItem{
	ControllerItem{
		Url:    "authorize",
		Method: "Get",
		HandlerFunc: func(context echo.Context) error {
			return context.JSON(http.StatusOK, "authorize")
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

func NewOauthController(e *echo.Echo, validatorLocate *NewMiddleware.ValidatorLocate) *EntityController {
	entityCtrl := new(EntityController)
	entityCtrl.intialEntityController(e, "oauth", validatorLocate)
	for _, ctrlItem := range privateCotrollerItem {
		entityCtrl.AddCtrlItem(ctrlItem)
	}
	return entityCtrl
}