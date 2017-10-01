package Model

import "github.com/lib/pq"

type QueryOauth struct {
	Entity
	ClientId     string
	CallBack     string
	ResponseType string
	Scope        pq.StringArray `gorm:"type:varchar(64)[]"`
	State        string
	Query        string
}
