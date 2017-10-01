package Dtos

type IdentityOAuthWithUsernameDto struct {
	TypeOAuth string `json:"type_oauth"  validate:"required"`
	QueryCode string `json:"query_code"  validate:"required"`
	UserName  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required"`
}
