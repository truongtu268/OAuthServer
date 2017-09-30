package Dtos

import "github.com/truongtu268/OAuthServer/Model"

type EntityDto interface {
	MapperDto2Entity() Model.IEntity
}
