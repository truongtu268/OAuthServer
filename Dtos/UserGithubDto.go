package Dtos

import (
	"github.com/truongtu268/OAuthServer/Model"
	"strconv"
)

type UserGithubDto struct {
	Id        int    `json:"id"`
	AvatarUrl string `json:"avatarUrl"`
	Name      string `json:"name"`
	Login     string `json:"login"`
	Email     string `json:"email"`
}

func (github *UserGithubDto) MapperDto2Entity() Model.IEntity {
	user := new(Model.User)
	user.Avatar = github.AvatarUrl
	user.Email = github.Email
	user.Name = github.Name
	user.SecurityInfos = []Model.UserSecurityInfo{
		Model.UserSecurityInfo{
			Username:       github.Login,
			IdFromProvider: strconv.Itoa(github.Id),
		},
	}
	return user
}
