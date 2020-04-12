package model

import (
	"context"

	"github.com/cymonevo/secret-api/entity"
	"github.com/cymonevo/secret-api/internal/log"
	"github.com/cymonevo/secret-api/internal/util"
	"github.com/cymonevo/secret-api/internal/validator"
	admin "github.com/cymonevo/secret-api/module/admin/repo"
	"github.com/cymonevo/secret-api/module/secret/repo"
)

const getSecretTag = "Secret|Get"

type GetSecretModel struct {
	adminRepo admin.DBRepo
	dbRepo    repo.DBRepo
	request   entity.GetSecretRequest
}

func (m *GetSecretModel) Do(ctx context.Context) (entity.GetSecretResponse, error) {
	var response entity.GetSecretResponse
	if err := m.Validate(ctx); err != nil {
		log.ErrorDetail(getSecretTag, "error validation: %v", err)
		response.Message = err.Error()
		return response, err
	}
	app, err := m.adminRepo.GetApp(ctx, m.request.AppID)
	if err != nil {
		log.ErrorDetail(insertSecretTag, "error get app data: %v", err)
		response.Message = err.Error()
		return response, err
	}
	var key [32]byte
	copy(key[:], app.Secret)
	last, err := m.dbRepo.GetLastSecret(ctx, m.request.AppID)
	if err != nil {
		log.ErrorDetail(insertSecretTag, "error get last data: %v", err)
		response.Message = err.Error()
		return response, err
	}
	old, err := util.Decrypt(last.Data, &key)
	if err != nil {
		log.ErrorDetail(insertSecretTag, "error decrypt old data: %v", err)
		response.Message = err.Error()
		return response, err
	}
	response.Data = old
	return response, nil
}

func (m *GetSecretModel) Validate(_ context.Context) error {
	v := validator.New()
	if m.request.AppID == "" {
		v.Missing("app_id")
	}
	return v.Error()
}
