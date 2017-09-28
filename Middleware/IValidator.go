package NewMiddleware

import "github.com/labstack/echo"

type IValidator interface {
	Execute(next echo.HandlerFunc) echo.HandlerFunc
}
