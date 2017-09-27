package Domain

import (
	"github.com/truongtu268/OAuthServer/Model"
)

type UserRepo struct {
	EntityRepository
}

func (userRepo *UserRepo) FindOrCreateUserByProviderLogin(user *Model.User) error {
	dbResult := userRepo.db.Where(Model.User{
		ProviderLogin:  user.ProviderLogin,
		IdFromProvider: user.IdFromProvider,
	}).FirstOrCreate(&user)
	return dbResult.Error
}
