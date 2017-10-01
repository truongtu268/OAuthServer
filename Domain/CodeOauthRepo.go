package Domain

import "github.com/truongtu268/OAuthServer/Model"

type CodeOauthRepo struct {
	EntityRepository
}

func (repo *CodeOauthRepo) FindQueryByQuery(codeOAuth string, query *Model.CodeOauth) error {
	dbResult := repo.db.Where(Model.CodeOauth{
		CodeOauth: codeOAuth,
	}).First(&query)
	return dbResult.Error
}

func NewCodeOauthRepo() *CodeOauthRepo {
	var repo = new(CodeOauthRepo)
	repo.InitialRepo(new(Model.CodeOauth), "")
	return repo
}
