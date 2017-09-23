package Controller

import (
	"github.com/labstack/echo"
	"fmt"
)

type EntityController struct {
	echo *echo.Echo
	subUrl string
	listEntityCtrItem []ControllerItem
}

func (entityCtr *EntityController)intialEntityController(e *echo.Echo, subUrl string)  {
	entityCtr.echo = e
	entityCtr.subUrl = subUrl
}

func (entityCtr *EntityController)AddCtrlItem(ctrItem ControllerItem)  {
	entityCtr.listEntityCtrItem = append(entityCtr.listEntityCtrItem, ctrItem)
}

func (entityCtr *EntityController)Execute() error {
	for _, ctrItem := range entityCtr.listEntityCtrItem {
		switch ctrItem.Method {
		case "Post": entityCtr.echo.POST(fmt.Sprintf("%s%s",entityCtr.subUrl,ctrItem.Url),ctrItem.HandlerFunc)
		case "Get": entityCtr.echo.GET(fmt.Sprintf("%s%s",entityCtr.subUrl,ctrItem.Url),ctrItem.HandlerFunc)
		case "Put": entityCtr.echo.PUT(fmt.Sprintf("%s%s",entityCtr.subUrl,ctrItem.Url),ctrItem.HandlerFunc)
		case "Delete": entityCtr.echo.DELETE(fmt.Sprintf("%s%s",entityCtr.subUrl,ctrItem.Url),ctrItem.HandlerFunc)
		}
	}
	return nil
}

func NewEntityController(e *echo.Echo, subUrl string) *EntityController {
	entityCtrl := new(EntityController)
	entityCtrl.intialEntityController(e,subUrl)
	return entityCtrl
}