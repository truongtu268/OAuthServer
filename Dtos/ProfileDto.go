package Dtos

import "time"

type ProfileDto struct {
	Avatar  string `json:"avatar" validate:"required"`
	Country string `json:"country" validate:"required"`
	DateOfBirth time.Time `json:"dob"`
}