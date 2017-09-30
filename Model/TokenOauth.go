package Model

import "time"

type TokenOauth struct {
	Entity
	AccessToken   string
	RefeshToken   string
	Expiry        time.Time
	TokenType     string
	Scope         string
	Provider      string
	UserRefer     string
	SecurityLevel string
}
