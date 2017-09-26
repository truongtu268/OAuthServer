package Domain

import (
	"io/ioutil"
	"encoding/json"
	"log"
)

type EndPoint struct {
	AuthURL string `json:"authurl" validate:"required"`
	TokenURL string `json:"tokenurl" validate:"required"`
}

type Provider struct {
	Name string `json:"name" validate:"required"`
	Cid string `json:"cid" validate:"required"`
	Csecret string `json:"csecret" validate:"required"`
	Callback string `json:"callback" validate:"required"`
	Scope []string `json:"scope" validate:"required"`
	Client string `json:"client" validate:"required"`
	EndPoint EndPoint `json:"endpoint" validate:"required"`
}

type DataConfig struct {
	Host string `json:"host" validate:"required"`
	User string `json:"user" validate:"required"`
	DbName string `json:"dbname" validate:"required"`
	Password string `json:"password" validate:"required"`
	Migrate string `json:"migrate" validate:"required"`
	OauthProviders []Provider `json:"oauth_providers"`
}

func LoadConfig(path string) DataConfig {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Config File Missing. ", err)
	}

	var config DataConfig
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatal("Config Parse Error: ", err)
	}

	return config
}

func GetConfigFile() chan DataConfig {
	channel:= make(chan DataConfig)
	go func() {
		channel<- LoadConfig("config.json")
	}()
	return channel
}