package Domain

import (
	"github.com/truongtu268/real_project/Model"
	"github.com/truongtu268/real_project/FakeData"
	"github.com/truongtu268/real_project/Common"
	"reflect"
	"strings"
)

type UnitOfWork struct {
	Config DataConfig
	Repositories map[string]*EntityRepository
}

func (unit*UnitOfWork) AddRepo(repo *EntityRepository, entity Model.IEntity) {
	var NameRepo = strings.Replace(reflect.TypeOf(entity).String(),"*Model.","Repo",-1)
	unit.Repositories[NameRepo] = repo
}

func (unit *UnitOfWork) Run() {
	unit.Repositories = make(map[string]*EntityRepository)
	unit.Config = <- GetConfigFile()
	if unit.Config.Migrate == "drop"{
		entitylist := []Model.IEntity{}
		entitylist = append(entitylist,
			new(Model.User),
			new(Model.PaymentMethod),
			new(Model.Store),
			new(Model.StatusOfShipping),
			new(Model.ProductCategory),
			new(Model.Inventory),
			new(Model.Address),
			new(Model.WareHouse),
			new(Model.SecurityInfo),
			new(Model.Profile),
			new(Model.TypeRatingCategory),
			new(Model.Rating),
			new(Model.LanguageContentEntity),
			new(Model.BillList),
			new(Model.Comment),
			new(Model.Reaction),
			new(Model.PaymentInfo),
			new(Model.ShippingInfo),
			new(Model.StatusOfBill),
			new(Model.BillItem))
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
	userRepo :=unit.Repositories["RepoUser"]
	storeRepo :=unit.Repositories["RepoStore"]
	paymentRepo :=unit.Repositories["RepoPaymentMethod"]
	for _, value := range FakeData.Users {
		var user = new(Model.User)
		Common.MapObject(value,user)
		userRepo.Create(user)
	}
	for _, value := range FakeData.Stores {
		var user = new(Model.Store)
		Common.MapObject(value,user)
		storeRepo.Create(user)
	}
	for _, value := range FakeData.PaymentMethod {
		var user = new(Model.PaymentMethod)
		Common.MapObject(value,user)
		paymentRepo.Create(user)
	}
}
