package Model

import (
	"time"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	uid "github.com/pborman/uuid"
	"encoding/base64"
)

type TokenOauth struct {
	ID            string `gorm:"primary_key;"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time
	AccessToken   string
	RefreshToken  string
	Expiry        time.Time
	TokenType     string
	Scope         string
	UserInfo      User   `gorm:"ForeignKey:UserTokRefer"`
	UserTokRefer  string
	SecurityLevel string
}

func (entity *TokenOauth) BeforeCreate(scope *gorm.Scope) error {
	accesstoken, refreshtoken, _ := entity.GenerateAccessToken(true)
	scope.SetColumn("ID", uuid.NewV4().String())
	scope.SetColumn("CreatedAt", time.Now())
	scope.SetColumn("UpdatedAt", time.Now())
	scope.SetColumn("AccessToken", accesstoken)
	scope.SetColumn("RefreshToken", refreshtoken)
	scope.SetColumn("Expiry", time.Now().Add(time.Duration(3600)*time.Second))
	return nil
}

func (entity *TokenOauth) GenerateAccessToken(generaterefresh bool) (accessToken string, refeshToken string, err error) {
	token := uid.NewRandom()
	accessToken = base64.RawURLEncoding.EncodeToString([]byte(token))
	if generaterefresh {
		rtoken := uid.NewRandom()
		refeshToken = base64.RawURLEncoding.EncodeToString([]byte(rtoken))
	}
	return
}

func (entity *TokenOauth) IsExpiredAt(t time.Time) bool {
	return entity.Expiry.Before(t)
}

func (entity *TokenOauth) IsExpired() bool {
	return entity.IsExpiredAt(time.Now())
}

func (user *TokenOauth) BeforeUpdate(scope gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}
