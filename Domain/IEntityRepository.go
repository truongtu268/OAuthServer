package Domain

import "github.com/truongtu268/OAuthServer/Model"

type Repository interface {
	Create(user Model.IEntity) error
	InitialTable() error
	InitialRepo(entity Model.IEntity, migrate string)
}
