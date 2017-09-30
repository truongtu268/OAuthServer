package Controller

import (
	"github.com/labstack/echo"
	"fmt"
	"github.com/truongtu268/OAuthServer/Middleware"
	"strings"
)

type EntityController struct {
	echo              *echo.Echo
	subUrl            string
	listEntityCtrItem []ControllerItem
	validateLocation  *NewMiddleware.ValidatorLocate
}

func (entityCtr *EntityController) intialEntityController(e *echo.Echo, subUrl string, validateLocation *NewMiddleware.ValidatorLocate) {
	entityCtr.echo = e
	entityCtr.subUrl = subUrl
	entityCtr.validateLocation = validateLocation
}

func (entityCtr *EntityController) AddCtrlItem(ctrItem ControllerItem) {
	entityCtr.listEntityCtrItem = append(entityCtr.listEntityCtrItem, ctrItem)
}

func (entityCtr *EntityController) Execute() error {
	for _, ctrItem := range entityCtr.listEntityCtrItem {
		fullUrl := fmt.Sprintf("%s/%s", entityCtr.subUrl, ctrItem.Url)
		switch ctrItem.Method {
		case "Post":
			entityCtr.echo.POST(fullUrl, ctrItem.HandlerFunc, mapMiddlewareFromPolicies2CtrlItem(fullUrl, entityCtr.validateLocation)...)
		case "Get":
			entityCtr.echo.GET(fullUrl, ctrItem.HandlerFunc, mapMiddlewareFromPolicies2CtrlItem(fullUrl, entityCtr.validateLocation)...)
		case "Put":
			entityCtr.echo.PUT(fullUrl, ctrItem.HandlerFunc, mapMiddlewareFromPolicies2CtrlItem(fullUrl, entityCtr.validateLocation)...)
		case "Delete":
			entityCtr.echo.DELETE(fullUrl, ctrItem.HandlerFunc, mapMiddlewareFromPolicies2CtrlItem(fullUrl, entityCtr.validateLocation)...)
		}
	}
	return nil
}
func TranferListStringToListValidator(listString []string, locate *NewMiddleware.ValidatorLocate) []echo.MiddlewareFunc {
	var listValidator = []echo.MiddlewareFunc{}
	for _, validatorName := range listString {
		err, validate := locate.GetValidator(validatorName)
		if err != nil {
			continue
		}
		listValidator = append(listValidator, validate.Execute)
	}
	return listValidator
}

func mapMiddlewareFromPolicies2CtrlItem(fullUrl string, validateLocation *NewMiddleware.ValidatorLocate) []echo.MiddlewareFunc {
	listUrl := strings.Split(fullUrl, "/")
	_, ok := Policies[listUrl[0]]
	if ok {
		listValidatorName, ok1 := Policies[listUrl[0]][listUrl[1]]
		if ok1 {
			return TranferListStringToListValidator(listValidatorName, validateLocation)
		}
		listValidatorNameDefaultCtroller, ok2 := Policies[listUrl[0]]["*"]
		if ok2 {
			return TranferListStringToListValidator(listValidatorNameDefaultCtroller, validateLocation)
		}
		listValidatorNameDefaultSystem := Policies["*"]["*"]
		return TranferListStringToListValidator(listValidatorNameDefaultSystem, validateLocation)
	}
	listValidatorNameDefaultSystem := Policies["*"]["*"]
	return TranferListStringToListValidator(listValidatorNameDefaultSystem, validateLocation)
}
