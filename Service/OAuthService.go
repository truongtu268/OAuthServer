package Service

import (
	"github.com/labstack/echo"
	"github.com/truongtu268/OAuthServer/Domain"
	"fmt"
	"io/ioutil"
	"golang.org/x/oauth2"
	"net/http"
	"errors"
	"crypto/rand"
	"encoding/base64"
)


type IOAuthService interface {
	OAuthFunc(e echo.Context) error
	LoginFunc(e echo.Context) error
	InitialFunc(conf Domain.Provider)
}

func randToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

type ProviderAuth struct {
	Provider Domain.Provider
	Conf oauth2.Config
}

func (googleOAuth *ProviderAuth)OAuthFunc(context echo.Context) error {
	queryState := context.QueryParam("state")
	fmt.Println(queryState)
	code := context.QueryParam("code")
	tok, err := googleOAuth.Conf.Exchange(oauth2.NoContext, code)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.New("Some thing wrong with token"))
	}
	fmt.Println("token",tok.AccessToken)
	fmt.Println("refeshtoken",tok.RefreshToken)
	fmt.Println(tok.Expiry)
	fmt.Println(tok.TokenType)
	client := googleOAuth.Conf.Client(oauth2.NoContext, tok)
	userinfo, err := client.Get(googleOAuth.Provider.Client)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.New("Some thing wrong with userinfo"))
	}
	defer userinfo.Body.Close()
	data, _ := ioutil.ReadAll(userinfo.Body)
	return context.JSON(http.StatusOK,string(data))
}

func (googleOAuth *ProviderAuth) LoginFunc(context echo.Context) error {
	state := randToken()
	fmt.Println(state)
	link := googleOAuth.Conf.AuthCodeURL(state)
	return context.JSON(http.StatusOK, link)
}

func (googleOAuth *ProviderAuth) InitialFunc(config Domain.Provider) {
	googleOAuth.Provider = config
	googleOAuth.Conf = oauth2.Config{
		ClientID:     googleOAuth.Provider.Cid,
		ClientSecret: googleOAuth.Provider.Csecret,
		RedirectURL:  googleOAuth.Provider.Callback,
		Scopes: googleOAuth.Provider.Scope,
		Endpoint: oauth2.Endpoint{
			AuthURL: googleOAuth.Provider.EndPoint.AuthURL,
			TokenURL: googleOAuth.Provider.EndPoint.TokenURL,
		},
	}
}

type OAuthService struct {
	oauthService map[string]IOAuthService
}

func (service *OAuthService)AddService(provider IOAuthService, name string)  {
	service.oauthService[name] = provider
}

func (service *OAuthService)GetService(providerName string) IOAuthService {
	return service.oauthService[providerName]
}

func NewOAuthService(config Domain.DataConfig) *OAuthService {
	service := new(OAuthService)
	service.oauthService = make(map[string]IOAuthService)
	for _, providerConf := range config.OauthProviders {
		provider := new(ProviderAuth)
		provider.InitialFunc(providerConf)
		service.AddService(provider, providerConf.Name)
	}
	return service
}