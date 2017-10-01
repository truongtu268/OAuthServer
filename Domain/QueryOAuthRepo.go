package Domain

import "github.com/truongtu268/OAuthServer/Model"

type QueryOAuthRepo struct {
	EntityRepository
}

func (repo *QueryOAuthRepo) FindQueryByQuery(queryCode string, query *Model.QueryOauth) error {
	dbResult := repo.db.Where(Model.QueryOauth{
		Query: queryCode,
	}).First(&query)
	return dbResult.Error
}

func (repo *QueryOAuthRepo) DeleteQueryByQueryString(queryString string) error {
	dbResult := repo.db.Where(Model.QueryOauth{
		Query: queryString,
	}).Delete(Model.QueryOauth{})
	return dbResult.Error
}

func NewQueryOAuthRepo() *QueryOAuthRepo {
	var repo = new(QueryOAuthRepo)
	repo.InitialRepo(new(Model.QueryOauth), "")
	return repo
}
