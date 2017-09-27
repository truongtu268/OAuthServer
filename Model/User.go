package Model

import (
	"time"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID         string `gorm:"primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Avatar string
	Email string
	Name string
	ProviderLogin string
	IdFromProvider string
	Username string
	Password string
}

func (user *User)HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (user *User)CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (entity *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4().String())
	scope.SetColumn("CreatedAt", time.Now())
	scope.SetColumn("UpdatedAt", time.Now())
	if len(entity.Password) >0 {
		passHash,_ := entity.HashPassword(entity.Password)
		scope.SetColumn("Password", passHash)
	}
	return nil
}

func (user *User) BeforeUpdate(scope gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}
