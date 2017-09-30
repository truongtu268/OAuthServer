package Service

import (
	"github.com/mitchellh/mapstructure"
	"golang.org/x/oauth2"
	"io/ioutil"
	"encoding/json"
	"errors"
	"github.com/truongtu268/OAuthServer/Dtos"
	"github.com/truongtu268/OAuthServer/Domain"
	"github.com/truongtu268/OAuthServer/Common"
	"github.com/truongtu268/OAuthServer/Model"
)

func MapTokenDto2Entity(tok *oauth2.Token, providerName string, user *Model.User) *Model.TokenOauth {
	token := new(Model.TokenOauth)
	token.AccessToken = tok.AccessToken
	token.Expiry = tok.Expiry
	token.Provider = providerName
	token.RefeshToken = tok.RefreshToken
	token.TokenType = tok.TokenType
	token.UserRefer = user.ID
	return token
}

func createUserAndUserSecurityInfo(
	userDto Dtos.EntityDto,
	tok *oauth2.Token,
	auth *ProviderAuth) (error, *Dtos.UserDto) {
	var userMapper = userDto.MapperDto2Entity()
	userEntity := new(Model.User)
	Common.MapObject(userMapper, userEntity)
	userEntity.SecurityInfos[0].ClientId = auth.Provider.ID
	err := userRepo.FindOrCreateUserByProviderLogin(userEntity)
	if err != nil {
		return err, nil
	}
	tokenEntity := MapTokenDto2Entity(tok, auth.Provider.Name, userEntity)
	tokenRepo.FindOrCreateTokenByProviderLogin(tokenEntity)
	var dto = new(Dtos.UserDto)
	Common.MapObject(userEntity, dto)
	return nil, dto
}

type IOAuthStorageData interface {
	CreateDataUserAndTokenToDataBase(auth *ProviderAuth,
		code string,
		userRepo *Domain.UserRepo,
		tokenRepo *Domain.TokenOauthRepo) (error, *Dtos.UserDto)
}

type GoogleStorageData struct{}

func (google *GoogleStorageData) CreateDataUserAndTokenToDataBase(
	auth *ProviderAuth,
	code string,
	userRepo *Domain.UserRepo,
	tokenRepo *Domain.TokenOauthRepo) (error, *Dtos.UserDto) {
	var userDto = &Dtos.UserGoogleDto{}
	tok, err := auth.Conf.Exchange(oauth2.NoContext, code)
	client := auth.Conf.Client(oauth2.NoContext, tok)
	userinfo, err := client.Get(auth.Provider.Client)
	if err != nil {
		return err, nil
	}
	defer userinfo.Body.Close()
	data, _ := ioutil.ReadAll(userinfo.Body)
	err = json.Unmarshal([]byte(data), userDto)
	if err != nil {
		return err, nil
	}
	return createUserAndUserSecurityInfo(userDto, tok, auth)
}

type GithubStorageData struct{}

func (github *GithubStorageData) CreateDataUserAndTokenToDataBase(
	auth *ProviderAuth,
	code string,
	userRepo *Domain.UserRepo,
	tokenRepo *Domain.TokenOauthRepo) (error, *Dtos.UserDto) {
	var userDto = &Dtos.UserGithubDto{}
	tok, err := auth.Conf.Exchange(oauth2.NoContext, code)
	client := auth.Conf.Client(oauth2.NoContext, tok)
	userinfo, err := client.Get(auth.Provider.Client)
	if err != nil {
		return err, nil
	}
	defer userinfo.Body.Close()
	data, _ := ioutil.ReadAll(userinfo.Body)
	err = json.Unmarshal([]byte(data), &userDto)
	if err != nil {
		return err, nil
	}
	return createUserAndUserSecurityInfo(userDto, tok, auth)
}

type InstagramStorageData struct{}

func (instagram *InstagramStorageData) CreateDataUserAndTokenToDataBase(
	auth *ProviderAuth,
	code string,
	userRepo *Domain.UserRepo,
	tokenRepo *Domain.TokenOauthRepo) (error, *Dtos.UserDto) {
	var userDto = &Dtos.UserInstagramDto{}
	tok, err := auth.Conf.Exchange(oauth2.NoContext, code)
	if err != nil {
		return err, nil
	}
	mapstructure.Decode(tok.Extra("user"), &userDto)
	return createUserAndUserSecurityInfo(userDto, tok, auth)
}

type ServiceLocateForStorageUserAndToken struct {
	listIOAuthStorageData map[string]IOAuthStorageData
}

func (services *ServiceLocateForStorageUserAndToken) AddOAuthStorageData(service IOAuthStorageData, name string) error {
	_, ok := services.listIOAuthStorageData[name]
	if ok {
		return errors.New("This key exist in Service Locate")
	}
	services.listIOAuthStorageData[name] = service
	return nil
}

func (services *ServiceLocateForStorageUserAndToken) GetOAuthStorageData(name string) (error, IOAuthStorageData) {
	service, ok := services.listIOAuthStorageData[name]
	if ok {
		return nil, service
	}
	return errors.New("This service doesn't exist in Service locator"), nil
}

func NewServiceLocateForStorageData() *ServiceLocateForStorageUserAndToken {
	services := new(ServiceLocateForStorageUserAndToken)
	services.listIOAuthStorageData = make(map[string]IOAuthStorageData)
	services.AddOAuthStorageData(new(GoogleStorageData), "google")
	services.AddOAuthStorageData(new(GithubStorageData), "github")
	services.AddOAuthStorageData(new(InstagramStorageData), "instagram")
	return services
}
