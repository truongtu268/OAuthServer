package Dtos

type IdentityOAuthWithAccessTokenDto struct {
	TypeOAuth    string `json:"type_oauth"  validate:"required"`
	QueryCode    string `json:"query_code"  validate:"required"`
	AccessToken  string `json:"access_token"  validate:"required"`
	RefreshToken string `json:"refresh_token"  validate:"required"`
}
