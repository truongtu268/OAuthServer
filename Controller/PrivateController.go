package Controller

import (
	"github.com/labstack/echo"
	"github.com/truongtu268/real_project/Common"
	"github.com/truongtu268/OAuthServer/Model"
	"net/http"
)

var privateCotrollerItem = []ControllerItem{
	ControllerItem{
		Url:"info",
		Method:"Get",
		HandlerFunc: func(context echo.Context) error {
			user := context.Get("user_info")
			userShow := new(Model.User)
			Common.MapObject(user,userShow)
			return context.JSON(http.StatusOK,userShow)
		},
	},
	ControllerItem{
		Url:"transaction",
		Method:"Get",
		HandlerFunc: func(context echo.Context) error {
			user := context.Get("user_info")
			userShow := new(Model.User)
			Common.MapObject(user,userShow)
			return context.JSON(http.StatusOK,userShow)
		},
	},
}
func NewPrivateController(e *echo.Echo) *EntityController {
	entityCtrl := new(EntityController)
	entityCtrl.intialEntityController(e,"private")
	for _, ctrlItem := range privateCotrollerItem {
		entityCtrl.AddCtrlItem(ctrlItem)
	}
	return entityCtrl
}