package Dtos

import "github.com/truongtu268/OAuthServer/Model"

type UserGoogleDto struct {
	Sub string `json:"sub"`
	Name string `json:"name"`
	GivenName string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Picture string `json:"picture"`
	Email string `json:"email"`
}

func (github *UserGoogleDto)MapperDto2Entity() Model.IEntity {
	user:= new(Model.User)
	user.Avatar = github.Picture
	user.Email = github.Email
	user.Username = github.Email
	user.Name = github.Name
	user.IdFromProvider = github.Sub
	user.ProviderLogin = "google"
	return user
}