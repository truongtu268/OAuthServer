package NewMiddleware

import (
	"github.com/labstack/echo"
	"errors"
	"strings"
	"net/http"

	"github.com/truongtu268/OAuthServer/Domain"
	"github.com/truongtu268/OAuthServer/Model"
)
type ValidateHeaderToken struct {}

func (validator *ValidateHeaderToken)Execute(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		tokenRepo := new(Domain.TokenOauthRepo)
		tokenRepo.InitialRepo(new(Model.TokenOauth),"")
		tokens := context.Request().Header.Get("Authorization")
		if len(tokens)<=0 {
			return context.JSON(http.StatusUnauthorized,"Not found Authorization header")
		}
		tokenString := strings.TrimPrefix(tokens,"Bearer ")
		if tokenString == "" {
			return context.JSON(http.StatusUnauthorized,errors.New("Token string is empty"))
		}
		userFromToken:= new(Model.User)
		err:= tokenRepo.FindAccessTokenToValidateUser(tokenString,userFromToken)
		if err != nil {
			return context.JSON(http.StatusUnauthorized,err)
		}
		context.Set("user_info",userFromToken)
		return next(context)
	}
}

type ValidateTest struct {}

func (validator *ValidateTest)Execute(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		context.Set("Test","Test success")
		return next(context)
	}
}