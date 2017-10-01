package Domain

import "github.com/truongtu268/OAuthServer/Model"

type ProviderRepo struct {
	EntityRepository
}

func (userRepo *ProviderRepo) Find(users *[]Model.Provider) error {
	dbResult := userRepo.db.Find(&users)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}

func NewProviderRepo() *ProviderRepo {
	var repo = new(ProviderRepo)
	repo.InitialRepo(new(Model.Provider), "")
	return repo
}
