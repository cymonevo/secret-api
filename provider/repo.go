package provider

import (
	"sync"

	base "github.com/cymonevo/secret-api/internal/base/repo"
	secret "github.com/cymonevo/secret-api/module/secret/repo"
)

var (
	baseDBRepo     *base.BaseDBRepo
	syncBaseDBRepo sync.Once

	secretDBRepo     secret.DBRepo
	syncSecretDBRepo sync.Once
)

func GetBaseDBRepo() *base.BaseDBRepo {
	syncBaseDBRepo.Do(func() {
		baseDBRepo = base.NewBaseDBRepo(GetDBClient())
	})
	return baseDBRepo
}

func GetSecretDBRepo() secret.DBRepo {
	syncSecretDBRepo.Do(func() {
		secret.NewSecretDBRepo(GetBaseDBRepo())
	})
	return secretDBRepo
}
