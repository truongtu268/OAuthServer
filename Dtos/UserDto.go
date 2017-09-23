package Dtos

type UserDto struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Profile ProfileDto `json:"profile"`
}