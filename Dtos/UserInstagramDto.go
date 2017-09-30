package Dtos

import "github.com/truongtu268/OAuthServer/Model"

type UserInstagramDto struct {
	Bio            string `json:"bio"`
	FullName       string `json:"full_name"`
	Id             string `json:"id"`
	IsBusiness     bool   `json:"is_bussiness"`
	ProfilePicture string `json:"profile_picture"`
	Username       string `json:"username"`
	Website        string `json:"website"`
}

func (github *UserInstagramDto) MapperDto2Entity() Model.IEntity {
	user := new(Model.User)
	user.Avatar = github.ProfilePicture
	user.Name = github.FullName
	user.SecurityInfos = []Model.UserSecurityInfo{
		Model.UserSecurityInfo{
			Username:       github.Username,
			IdFromProvider: github.Id,
			ProviderLogin:  "instagram",
		},
	}
	return user
}
