package Dtos

import "time"

type TokenDto struct {
	AccessToken string `json:"access_token"`
	TokenType string `json:"token_type,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	Expiry time.Time `json:"expiry,omitempty"`
}