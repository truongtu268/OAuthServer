package Domain

import "github.com/truongtu268/OAuthServer/Model"

type ClientRepo struct {
	EntityRepository
}

func (clientRepo *ClientRepo) Find(clients *[]Model.Client) error {
	dbResult := clientRepo.db.Find(&clients)
	return dbResult.Error
}

func (clientRepo *ClientRepo) FindOne(id string, client *Model.Client) error {
	dbResult := clientRepo.db.Where(Model.Client{
		ID: id,
	}).First(&client)
	return dbResult.Error
}

func NewClientRepo() *ClientRepo {
	var repo = new(ClientRepo)
	repo.InitialRepo(new(Model.Client), "")
	return repo
}
