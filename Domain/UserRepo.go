package Domain

import (
	"github.com/truongtu268/OAuthServer/Model"
)

type UserRepo struct {
	EntityRepository
}

func (userRepo *UserRepo) FindOrCreateUserByProviderLogin(user *Model.User) error {
	securityInfo := new(Model.UserSecurityInfo)
	dbResult := userRepo.db.Where(Model.UserSecurityInfo{
		ProviderLogin:  user.SecurityInfos[0].ProviderLogin,
		IdFromProvider: user.SecurityInfos[0].IdFromProvider,
	}).First(&securityInfo)
	if dbResult.Error != nil {
		dbCreateUser := userRepo.db.Create(&user)
		return dbCreateUser.Error
	}
	dbResultToPopulateUser := userRepo.db.Model(user).Where(Model.User{
		ID: securityInfo.UserRefer,
	}).First(&user)
	return dbResultToPopulateUser.Error
}
