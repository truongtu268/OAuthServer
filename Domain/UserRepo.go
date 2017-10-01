package Domain

import (
	"github.com/truongtu268/OAuthServer/Model"
	"github.com/truongtu268/OAuthServer/Dtos"
	"errors"
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

func (userRepo *UserRepo) CheckSecurityInfo(user *Model.User, dto *Dtos.IdentityOAuthWithUsernameDto) error {
	securityInfo := new(Model.UserSecurityInfo)
	dbCheckUserNameResult := userRepo.db.Where(Model.UserSecurityInfo{
		Username: dto.UserName,
	}).First(&securityInfo)
	if dbCheckUserNameResult.Error != nil {
		return dbCheckUserNameResult.Error
	}
	if !securityInfo.CheckPasswordHash(dto.Password, securityInfo.Password) {
		return errors.New("Password is not match")
	}
	dbFindUserByIdResult := userRepo.db.Where(Model.User{
		ID: securityInfo.UserRefer,
	}).First(&user)
	return dbFindUserByIdResult.Error
}

func NewUserRepo() *UserRepo {
	var repo = new(UserRepo)
	repo.InitialRepo(new(Model.User), "")
	return repo
}
