package Dtos

import (
	"github.com/truongtu268/OAuthServer/Model"
	"fmt"
	"strconv"
)

type UserGithubDto struct {
	Id int `json:"id"`
	AvatarUrl string `json:"avatarUrl"`
	Name string `json:"name"`
	Login string `json:"login"`
	Email string `json:"email"`
}

func (github *UserGithubDto)MapperDto2Entity() Model.IEntity {
	user:= new(Model.User)
	user.Avatar = github.AvatarUrl
	user.Email = github.Email
	user.Username = github.Login
	user.Name = github.Name
	user.IdFromProvider = strconv.Itoa(github.Id)
	user.ProviderLogin = "github"
	fmt.Println(user)
	return user
}