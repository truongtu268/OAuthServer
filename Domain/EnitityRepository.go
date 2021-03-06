package Domain

import (
	"github.com/jinzhu/gorm"
	"github.com/truongtu268/OAuthServer/Model"
)

type EntityRepository struct {
	db          *gorm.DB
	entity      Model.IEntity
	migrateMode string
}

func (userRepo *EntityRepository) Create(user Model.IEntity) error {
	dbResult := userRepo.db.Create(user)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}

func (userRepo *EntityRepository) InitialRepo(entity Model.IEntity, migrate string) {
	userRepo.db = GetInstance()
	userRepo.entity = entity
	userRepo.migrateMode = migrate
}

func (userRepo *EntityRepository) InitialTable() error {
	if userRepo.migrateMode != "drop" {
		return nil
	}
	if userRepo.db.HasTable(userRepo.entity) {
		userRepo.db.DropTable(userRepo.entity)
	}
	userRepo.db.CreateTable(userRepo.entity)
	return nil
}
