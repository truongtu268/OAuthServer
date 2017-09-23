package Domain

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var instance *gorm.DB

func GetInstance() *gorm.DB {
	if instance ==nil {
		instance, _ = gorm.Open("postgres","host=localhost user=postgres dbname=testORM sslmode=disable password=truongtu268")
		return instance
	}
	return instance
}
