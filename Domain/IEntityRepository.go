package Domain

import "github.com/truongtu268/real_project/Model"

type Repository interface {
	Create(user Model.IEntity) error
	Delete(id string, entity Model.IEntity) error
	Find(entities [] Model.IEntity) error
	FindOne(id string, entity Model.IEntity) error
	InitialTable() error
	InitialRepo(entity Model.IEntity, migrate string)
}