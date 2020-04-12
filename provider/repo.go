package provider

import (
	"sync"

	"github.com/cymonevo/secret-api/internal/base/repo"
	admin "github.com/cymonevo/secret-api/module/admin/repo"
	secret "github.com/cymonevo/secret-api/module/secret/repo"
)

var (
	baseDBRepo     *repo.BaseDBRepo
	syncBaseDBRepo sync.Once

	adminDBRepo     admin.DBRepo
	syncAdminDBRepo sync.Once

	secretDBRepo     secret.DBRepo
	syncSecretDBRepo sync.Once
)

func GetBaseDBRepo() *repo.BaseDBRepo {
	syncBaseDBRepo.Do(func() {
		baseDBRepo = repo.NewBaseDBRepo(GetDBClient())
	})
	return baseDBRepo
}

func GetAdminDBRepo() admin.DBRepo {
	syncAdminDBRepo.Do(func() {
		adminDBRepo = admin.NewAdminDBRepo(GetBaseDBRepo())
	})
	return adminDBRepo
}

func GetSecretDBRepo() secret.DBRepo {
	syncSecretDBRepo.Do(func() {
		secretDBRepo = secret.NewSecretDBRepo(GetBaseDBRepo())
	})
	return secretDBRepo
}
