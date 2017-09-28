package Domain

import (
	"github.com/truongtu268/OAuthServer/Model"
)

type TokenOauthRepo struct {
	EntityRepository
}

func (userRepo *TokenOauthRepo) FindOrCreateTokenByProviderLogin(token *Model.TokenOauth) error {
	dbFindUserExist := userRepo.db.Where(Model.TokenOauth{
		UserRefer:token.UserRefer,
	}).Find(&token)
	if dbFindUserExist.Error != nil {
		dbCreateNewUserToken := userRepo.db.Create(&token)
		return dbCreateNewUserToken.Error
	}
	dbResultFindAccessTokenExist := userRepo.db.Where(Model.TokenOauth{
		AccessToken:token.AccessToken,
		Provider:token.Provider,
	}).Find(&token)
	if dbResultFindAccessTokenExist.Error !=nil {
		dbResultUpdateAccsessToken:= userRepo.db.Model(&token).Where("user_refer = ?",token.UserRefer).Update("access_token","expiry")
		return dbResultUpdateAccsessToken.Error
	}
	return nil
}