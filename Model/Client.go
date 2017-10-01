package Model

import (
	"time"
	"github.com/jinzhu/gorm"
	//"github.com/satori/go.uuid"
	"github.com/lib/pq"
)

type Client struct {
	ID         string         `gorm:"primary_key;"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
	Secret     string
	TrustLevel string
	CallBack   pq.StringArray `gorm:"type:varchar(64)[]"`
	Scope      pq.StringArray `gorm:"type:varchar(64)[]"`
	UserCreate User
}

func (entity *Client) BeforeCreate(scope *gorm.Scope) error {
	//scope.SetColumn("ID", uuid.NewV4().String())
	//scope.SetColumn("Secret", uuid.NewV4().String())
	scope.SetColumn("CreatedAt", time.Now())
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}

func (entity *Client) BeforeUpdate(scope gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}
