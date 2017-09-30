package Model

import (
	"time"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

type User struct {
	ID            string             `gorm:"primary_key;"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time
	Avatar        string
	Email         string
	Name          string
	SecurityInfos []UserSecurityInfo `gorm:"ForeignKey:UserRefer"`
	ClientRefer   string
}

func (entity *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4().String())
	scope.SetColumn("CreatedAt", time.Now())
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}

func (user *User) BeforeUpdate(scope gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}
