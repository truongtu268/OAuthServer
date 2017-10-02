package Service

import (
	"github.com/truongtu268/OAuthServer/Domain"
	"github.com/truongtu268/OAuthServer/Dtos"
	"github.com/truongtu268/OAuthServer/Model"
)

type UserTokenService struct {
	tokenRepo *Domain.TokenOauthRepo
	userRepo  *Domain.UserRepo
}

func (service *UserTokenService) CheckAccessTokenFromOAuthService(dto *Dtos.IdentityOAuthWithAccessTokenDto) (*Model.User, error) {
	user := new(Model.User)
	err := tokenRepo.FindAccessTokenToValidateUser(dto.AccessToken, user)
	if err != nil {
		switch err.Error() {
		case "Token is expired":
			tokenReGenerate := new(Model.TokenOauth)
			err = service.tokenRepo.RefreshAccessToken(dto.RefreshToken, tokenReGenerate)
			if err != nil {
				return nil, err
			}
			err = service.tokenRepo.FindAccessTokenToValidateUser(tokenReGenerate.AccessToken, user)
			if err != nil {
				return nil, err
			}
			return user, nil
		default:
			return nil, err
		}
	}
	return user, nil
}

func (service *UserTokenService) CheckUserSecurityInfo(dto *Dtos.IdentityOAuthWithUsernameDto) (*Model.User, error) {
	user := new(Model.User)
	err := service.userRepo.CheckSecurityInfo(user, dto)
	if err != nil {
		return nil, err
	}
	token := new(Model.TokenOauth)
	token.UserTokRefer = user.ID
	token.TokenType = "code"
	service.tokenRepo.Create(token)
	return user, nil
}

func NewUserTokenService() *UserTokenService {
	service := new(UserTokenService)
	service.tokenRepo = Domain.NewTokenOauthRepo()
	service.userRepo = Domain.NewUserRepo()
	return service
}
