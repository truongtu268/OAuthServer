package Controller

import "github.com/labstack/echo"

type ControllerItem struct {
	Url string
	Method string
	HandlerFunc func(context echo.Context) error
}
