package Dtos

import "time"

type UserDto struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Avatar    string
	Email     string
	Name      string
}
