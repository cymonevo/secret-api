package model

import (
	"github.com/cymonevo/secret-api/entity"
	admin "github.com/cymonevo/secret-api/module/admin/repo"
	"github.com/cymonevo/secret-api/module/secret/repo"
)

type Factory struct {
	adminRepo admin.DBRepo
	dbRepo    repo.DBRepo
}

func NewSecretFactory(adminRepo admin.DBRepo, dbRepo repo.DBRepo) *Factory {
	return &Factory{
		adminRepo: adminRepo,
		dbRepo:    dbRepo,
	}
}

func (f *Factory) NewInsertSecretModel(request entity.InsertSecretRequest) *InsertSecretModel {
	return &InsertSecretModel{
		adminRepo: f.adminRepo,
		dbRepo:    f.dbRepo,
		request:   request,
	}
}
