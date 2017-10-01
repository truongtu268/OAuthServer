package Dtos

type IdentityOAuthDto struct {
	TypeOAuth    string `json:"type_oauth"  validate:"required"`
	QueryCode    string `json:"query_code"  validate:"required"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	UserName     string `json:"username"`
	Password     string `json:"password"`
}
