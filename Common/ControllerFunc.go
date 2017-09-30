package Common

import "github.com/labstack/echo"

type ControllerFunc interface {
	Initial(urlParams map[string]*interface{}, dtos []*interface{}, context echo.Context)
	Execute() error
}
