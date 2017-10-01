package Service

import (
	"golang.org/x/oauth2"
	"github.com/truongtu268/OAuthServer/Domain"
	"github.com/labstack/echo"
	"github.com/truongtu268/OAuthServer/Model"
	"net/http"
	"github.com/truongtu268/OAuthServer/Common"
)

var serviceLocator *StorageUserAndTokenFromProviderService
var userRepo *Domain.UserRepo
var tokenRepo *Domain.TokenOauthRepo

type IProviderService interface {
	OAuthFunc(e echo.Context) error
	LoginFunc(e echo.Context) error
	InitialFunc(provider Model.Provider)
}

type ProviderAuth struct {
	Provider Model.Provider
	Conf     oauth2.Config
}

func (googleOAuth *ProviderAuth) OAuthFunc(context echo.Context) error {
	code := context.QueryParam("code")
	err, serviceStorage := serviceLocator.GetOAuthStorageData(googleOAuth.Provider.Name)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}
	err, user := serviceStorage.CreateDataUserAndTokenToDataBase(googleOAuth, code, userRepo, tokenRepo)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}
	return context.JSON(http.StatusOK, user)
}

func (googleOAuth *ProviderAuth) LoginFunc(context echo.Context) error {
	state := Common.RandToken()
	link := googleOAuth.Conf.AuthCodeURL(state)
	return context.JSON(http.StatusOK, link)
}

func (googleOAuth *ProviderAuth) InitialFunc(provider Model.Provider) {
	googleOAuth.Provider = provider
	googleOAuth.Conf = oauth2.Config{
		ClientID:     googleOAuth.Provider.Cid,
		ClientSecret: googleOAuth.Provider.Csecret,
		RedirectURL:  googleOAuth.Provider.Callback,
		Scopes:       googleOAuth.Provider.Scope,
		Endpoint: oauth2.Endpoint{
			AuthURL:  googleOAuth.Provider.AuthURL,
			TokenURL: googleOAuth.Provider.TokenURL,
		},
	}
}

type ProviderService struct {
	oauthService map[string]IProviderService
}

func (service *ProviderService) AddService(provider IProviderService, name string) {
	service.oauthService[name] = provider
}

func (service *ProviderService) GetService(providerName string) IProviderService {
	return service.oauthService[providerName]
}

func NewProviderService() *ProviderService {
	providerRepo := Domain.NewProviderRepo()
	userRepo = Domain.NewUserRepo()
	tokenRepo = Domain.NewTokenOauthRepo()
	serviceLocator = NewStorageUserAndTokenFromProviderService()
	service := new(ProviderService)
	service.oauthService = make(map[string]IProviderService)
	providerList := new([]Model.Provider)
	providerRepo.Find(providerList)
	for _, providerConf := range *providerList {
		provider := new(ProviderAuth)
		provider.InitialFunc(providerConf)
		service.AddService(provider, providerConf.Name)
	}
	return service
}
