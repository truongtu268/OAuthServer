package Domain

import (
	"github.com/truongtu268/OAuthServer/Model"
)

type TokenOauthRepo struct {
	EntityRepository
}

func (userRepo *TokenOauthRepo) FindOrCreateTokenByProviderLogin(token *Model.TokenOauth) error {
	dbResult := userRepo.db.Where(Model.TokenOauth{
		AccessToken:token.AccessToken,
	}).FirstOrCreate(&token)
	return dbResult.Error
}