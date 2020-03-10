package model

import (
	"github.com/cymonevo/secret-api/entity"
	"github.com/cymonevo/secret-api/module/admin/repo"
)

type Factory struct {
	dbRepo repo.DBRepo
}

func NewAdminFactory(dbRpo repo.DBRepo) *Factory {
	return &Factory{
		dbRepo: dbRpo,
	}
}

func (f *Factory) NewRegisterModel(request entity.RegisterRequest) *RegisterModel {
	return &RegisterModel{
		dbRepo:  f.dbRepo,
		request: request,
	}
}
