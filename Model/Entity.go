package Model

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"time"
)

type Entity struct {
	ID         string `gorm:"primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (entity *Entity) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4().String())
	scope.SetColumn("CreatedAt", time.Now())
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}

func (user *Entity) BeforeUpdate(scope gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}

