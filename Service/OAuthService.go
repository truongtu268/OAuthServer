package Service

import (
	"github.com/labstack/echo"
	"net/http"
	"encoding/base64"
	"github.com/truongtu268/OAuthServer/Model"
	"github.com/truongtu268/OAuthServer/Domain"
	"github.com/truongtu268/OAuthServer/Dtos"
	"github.com/truongtu268/OAuthServer/Common"
	"github.com/pborman/uuid"
	"fmt"
)

type OAuthService struct {
	clientRepo    *Domain.ClientRepo
	queryRepo     *Domain.QueryOAuthRepo
	tokenService  *UserTokenService
	codeOauthRepo *Domain.CodeOauthRepo
}

func (oauth *OAuthService) OAuthClientService(context echo.Context) error {
	var queryOauthCode = base64.RawURLEncoding.EncodeToString([]byte(uuid.NewRandom()))
	queryOauth := new(Dtos.QueryOAuthDto)
	err := context.Bind(queryOauth)
	if err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}
	err = context.Validate(queryOauth)
	if err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}
	client := new(Model.Client)
	oauth.clientRepo.FindOne(queryOauth.ClientId, client)
	if !Common.StringContains(client.CallBack, queryOauth.CallBack) {
		return context.JSON(http.StatusBadRequest, "Callback url not map")
	}
	if queryOauth.ResponseType != "code" {
		return context.JSON(http.StatusBadRequest, "ResponseType not right")
	}
	queryEntity := new(Model.QueryOauth)
	Common.MapObject(queryOauth, queryEntity)
	queryEntity.Query = queryOauthCode
	queryEntity.ClientId = client.ID
	queryEntity.ResponseType = "code"
	err = oauth.queryRepo.Create(queryEntity)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}
	return context.JSON(http.StatusOK, queryEntity.Query)
}

func (oauth *OAuthService) OAuthUserService(context echo.Context) error {
	identityDto := new(Dtos.IdentityOAuthDto)
	err := context.Bind(identityDto)
	if err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())
	}
	err = context.Validate(identityDto)
	if err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())
	}
	query := new(Model.QueryOauth)
	err = oauth.queryRepo.FindQueryByQuery(identityDto.QueryCode, query)
	if err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())
	}
	oauth.queryRepo.DeleteQueryByQueryString(query.Query)
	codeOAuth := new(Model.CodeOauth)
	Common.MapObject(query, codeOAuth)
	switch identityDto.TypeOAuth {
	case "accesstoken":
		accessTokenDto := new(Dtos.IdentityOAuthWithAccessTokenDto)
		Common.MapObject(identityDto, accessTokenDto)
		user, err := oauth.tokenService.CheckAccessTokenFromOAuthService(accessTokenDto)
		if err != nil {
			return context.JSON(http.StatusInternalServerError, err.Error())
		}
		codeOAuth.UserId = user.ID
		codeOAuth.CodeOauth = Common.RandToken()
		oauth.codeOauthRepo.Create(codeOAuth)
		return context.JSON(http.StatusOK, fmt.Sprintf("%s?state=%scode=%s", codeOAuth.CallBack, codeOAuth.State, codeOAuth.CodeOauth))
	case "username":
		userNameDto := new(Dtos.IdentityOAuthWithUsernameDto)
		Common.MapObject(identityDto, userNameDto)
		user, err := oauth.tokenService.CheckUserSecurityInfo(userNameDto)
		if err != nil {
			return context.JSON(http.StatusInternalServerError, err.Error())
		}
		codeOAuth.UserId = user.ID
		codeOAuth.CodeOauth = Common.RandToken()
		oauth.codeOauthRepo.Create(codeOAuth)
		return context.JSON(http.StatusOK, fmt.Sprintf("%s?state=%s&code=%s", codeOAuth.CallBack, codeOAuth.State, codeOAuth.CodeOauth))
	default:
		return context.JSON(http.StatusBadRequest, "invalide type oauth")
	}
}

func NewOAuthService() *OAuthService {
	service := new(OAuthService)
	service.queryRepo = Domain.NewQueryOAuthRepo()
	service.clientRepo = Domain.NewClientRepo()
	service.tokenService = NewUserTokenService()
	service.codeOauthRepo = Domain.NewCodeOauthRepo()
	return service
}
