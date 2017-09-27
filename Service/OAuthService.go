package Service

import (
	"github.com/labstack/echo"
	"golang.org/x/oauth2"
	"net/http"
	"crypto/rand"
	"encoding/base64"
	"github.com/truongtu268/OAuthServer/Model"
	"github.com/truongtu268/OAuthServer/Domain"
)

func randToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

var serviceLocator *ServiceLocateForStorageUserAndToken
var userRepo *Domain.UserRepo
var tokenRepo *Domain.TokenOauthRepo

type IOAuthService interface {
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
	err,serviceStorage := serviceLocator.GetOAuthStorageData(googleOAuth.Provider.Name)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}
	err,user :=serviceStorage.CreateDataUserAndTokenToDataBase(googleOAuth,code,userRepo, tokenRepo)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}
	return context.JSON(http.StatusOK, user)
}

func (googleOAuth *ProviderAuth) LoginFunc(context echo.Context) error {
	state := randToken()
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

type OAuthService struct {
	oauthService map[string]IOAuthService
}

func (service *OAuthService) AddService(provider IOAuthService, name string) {
	service.oauthService[name] = provider
}

func (service *OAuthService) GetService(providerName string) IOAuthService {
	return service.oauthService[providerName]
}

func NewOAuthService() *OAuthService {
	providerRepo := new(Domain.ProviderRepo)
	userRepo = new(Domain.UserRepo)
	tokenRepo = new(Domain.TokenOauthRepo)
	providerRepo.InitialRepo(new(Model.Provider), "")
	userRepo.InitialRepo(new(Model.User),"")
	tokenRepo.InitialRepo(new(Model.TokenOauth),"")
	serviceLocator = NewServiceLocateForStorageData()
	service := new(OAuthService)
	service.oauthService = make(map[string]IOAuthService)
	providerList := new([]Model.Provider)
	providerRepo.Find(providerList)
	for _, providerConf := range *providerList {
		provider := new(ProviderAuth)
		provider.InitialFunc(providerConf)
		service.AddService(provider, providerConf.Name)
	}
	return service
}
