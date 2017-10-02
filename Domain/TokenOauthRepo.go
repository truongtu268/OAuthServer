package Domain

import (
	"github.com/truongtu268/OAuthServer/Model"
	"errors"
)

type TokenOauthRepo struct {
	EntityRepository
}

func (userRepo *TokenOauthRepo) FindOrCreateTokenByProviderLogin(token *Model.TokenOauth) error {
	dbFindUserExist := userRepo.db.Where(Model.TokenOauth{
		UserTokRefer: token.UserTokRefer,
	}).Find(&token)
	if dbFindUserExist.Error != nil {
		dbCreateNewUserToken := userRepo.db.Create(&token)
		return dbCreateNewUserToken.Error
	}
	return dbFindUserExist.Error
}

func (tokenRepo *TokenOauthRepo) FindAccessTokenToValidateUser(token string, user *Model.User) error {
	tokenResult := new(Model.TokenOauth)
	dbFindAccessToken := tokenRepo.db.Where(Model.TokenOauth{
		AccessToken: token,
	}).First(&tokenResult)
	if dbFindAccessToken.Error != nil {
		return errors.New("Token doesn't validate")
	}
	if tokenResult.IsExpired() {
		return errors.New("Token is expired")
	}
	dbUserFindFromTok := tokenRepo.db.Where(Model.User{
		ID: tokenResult.UserTokRefer,
	}).First(&user)
	return dbUserFindFromTok.Error
}

func (tokenRepo *TokenOauthRepo) RefreshAccessToken(refreshToken string, token *Model.TokenOauth) error {
	dbFindRefreshToken := tokenRepo.db.Where(Model.TokenOauth{
		RefreshToken: refreshToken,
	}).First(&token)
	if dbFindRefreshToken.Error != nil {
		return errors.New("Refresh Token doesn't validate")
	}
	newAccessToken, _, err := token.GenerateAccessToken(false)
	if err != nil {
		return errors.New("Something happend with server")
	}
	dbUpdateTokenResult := tokenRepo.db.Model(&token).Update("access_token", newAccessToken)
	return dbUpdateTokenResult.Error
}

func (repo *TokenOauthRepo)FindTokenByUserId(userId string, token *Model.TokenOauth) error {
	dbFindTokenByUserIdResult := repo.db.Where(Model.TokenOauth{
		UserTokRefer:userId,
	}).First(&token)
	return dbFindTokenByUserIdResult.Error
}

func NewTokenOauthRepo() *TokenOauthRepo {
	var repo = new(TokenOauthRepo)
	repo.InitialRepo(new(Model.TokenOauth), "")
	return repo
}
