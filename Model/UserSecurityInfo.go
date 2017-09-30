package Model

import (
	"time"
	"golang.org/x/crypto/bcrypt"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

type UserSecurityInfo struct {
	ID             string `gorm:"primary_key;"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
	ClientId       string
	IdFromProvider string
	Username       string
	Password       string
	UserInfo       User   `gorm:"ForeignKey:UserRefer"`
	UserRefer      string
}

func (user *UserSecurityInfo) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (user *UserSecurityInfo) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (entity *UserSecurityInfo) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4().String())
	scope.SetColumn("CreatedAt", time.Now())
	scope.SetColumn("UpdatedAt", time.Now())
	if len(entity.Password) > 0 {
		passHash, _ := entity.HashPassword(entity.Password)
		scope.SetColumn("Password", passHash)
	}
	return nil
}

func (user *UserSecurityInfo) BeforeUpdate(scope gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}
