package Domain

import (
	"github.com/truongtu268/OAuthServer/Model"
	"reflect"
	"strings"
	"github.com/truongtu268/OAuthServer/FakeData"
	"github.com/truongtu268/OAuthServer/Common"
)

type UnitOfWork struct {
	Config DataConfig
	Repositories map[string]*EntityRepository
}

func (unit*UnitOfWork) AddRepo(repo *EntityRepository, entity Model.IEntity) {
	var NameRepo = strings.Replace(reflect.TypeOf(entity).String(),"*Model.","Repo",-1)
	unit.Repositories[NameRepo] = repo
}
var Config DataConfig
func (unit *UnitOfWork) Run() {
	unit.Repositories = make(map[string]*EntityRepository)
	Config = unit.Config
	if unit.Config.Migrate == "drop"{
		entitylist := []Model.IEntity{}
		entitylist = append(entitylist,
			new(Model.Provider),
			new(Model.User),
			new(Model.TokenOauth),)
		for _, entity := range entitylist {
			entityRepo := new(EntityRepository)
			entityRepo.InitialRepo(entity,unit.Config.Migrate)
			entityRepo.InitialTable()
			unit.AddRepo(entityRepo,entity)
		}
		unit.boostrapData()
	}
}

func (unit *UnitOfWork) boostrapData()  {
	providerRepo :=unit.Repositories["RepoProvider"]
	for _, value := range FakeData.Providers {
		var user = new(Model.Provider)
		Common.MapObject(value,user)
		providerRepo.Create(user)
	}
}
