package Model

import "github.com/jinzhu/gorm"

type IEntity interface {
	BeforeCreate(scope *gorm.Scope) error
}
