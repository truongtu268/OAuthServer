package Model

import "github.com/lib/pq"

type Provider struct {
	Entity
	Name     string
	Cid      string
	Csecret  string
	Callback string
	Scope    pq.StringArray `gorm:"type:varchar(64)[]"`
	Client   string
	AuthURL  string
	TokenURL string
}
