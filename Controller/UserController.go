package Controller

import (
	"github.com/labstack/echo"
	"github.com/truongtu268/OAuthServer/Domain"
	"github.com/truongtu268/OAuthServer/Model"
	"github.com/truongtu268/OAuthServer/Dtos"
	"net/http"
	"fmt"
	"github.com/truongtu268/OAuthServer/Common"
)
var UserCtrlItems = []ControllerItem{
	ControllerItem{
		Url:"",
		Method:"Post",
		HandlerFunc: func(c echo.Context) error {
				var repository = new(Domain.UserRepository)
				repository.InitialRepo(new(Model.User), "")
				dto := new(Dtos.UserDto)
				if err := c.Bind(dto); err != nil {
					return err
				}
			if err := c.Validate(dto); err != nil {
				return c.String(http.StatusBadRequest, fmt.Sprintf("%v", err))
			}
			u := new(Model.User)
			Common.MapObject(dto, u)
			err := repository.Create(u)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, err)
			}
			Common.MapObject(u, dto)
			return c.JSON(http.StatusOK, dto)
		},
	},
}

func NewUserController(e *echo.Echo) *EntityController {
	var userCTR = new(EntityController)
	userCTR.intialEntityController(e,"/users")
	for _, ctrlItem := range UserCtrlItems {
		userCTR.AddCtrlItem(ctrlItem)
	}
	return userCTR
}
