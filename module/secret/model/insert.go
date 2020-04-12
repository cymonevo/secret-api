package model

import (
	"context"

	"github.com/cymonevo/secret-api/entity"
	"github.com/cymonevo/secret-api/internal/encoding/json"
	"github.com/cymonevo/secret-api/internal/log"
	"github.com/cymonevo/secret-api/internal/util"
	"github.com/cymonevo/secret-api/internal/validator"
	admin "github.com/cymonevo/secret-api/module/admin/repo"
	"github.com/cymonevo/secret-api/module/secret/repo"
)

const insertSecretTag = "Secret|Insert"

type InsertSecretModel struct {
	adminRepo admin.DBRepo
	dbRepo    repo.DBRepo
	request   entity.InsertSecretRequest
}

func (m *InsertSecretModel) Do(ctx context.Context) (entity.InsertSecretResponse, error) {
	var response entity.InsertSecretResponse
	if err := m.Validate(ctx); err != nil {
		log.ErrorDetail(insertSecretTag, "error validation: %v", err)
		response.Message = err.Error()
		return response, err
	}
	app, err := m.adminRepo.GetApp(ctx, m.request.AppID)
	if err != nil {
		log.ErrorDetail(insertSecretTag, "error get app data: %v", err)
		response.Message = err.Error()
		return response, err
	}
	data, err := json.Marshal(m.request.Data)
	log.DebugDetail("[DEBUG] data: %v", data)
	var key [32]byte
	copy(key[:], app.Secret)
	last, err := m.dbRepo.GetLastSecret(ctx, m.request.AppID)
	if err != nil {
		if util.IsErrNotFound(err) {
			return m.insert(ctx, entity.SecretData{
				AppID: m.request.AppID,
				Data:  data,
			})
		}
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
	result, err := json.Merge(old, data, json.OptionReplace)
	if err != nil {
		log.ErrorDetail(insertSecretTag, "error merge data: %v", err)
		response.Message = err.Error()
		return response, err
	}
	last.Data, err = util.Encrypt(result, &key)
	if err != nil {
		log.ErrorDetail(insertSecretTag, "error encrypt result: %v", err)
		response.Message = err.Error()
		return response, err
	}
	return m.insert(ctx, last)
}

func (m *InsertSecretModel) insert(ctx context.Context, secret entity.SecretData) (resp entity.InsertSecretResponse, err error) {
	err = m.dbRepo.InsertSecret(ctx, secret)
	if err != nil {
		log.ErrorDetail(insertSecretTag, "error save to db: %v", err)
		resp.Message = err.Error()
		return
	}
	return
}

func (m *InsertSecretModel) Validate(_ context.Context) error {
	v := validator.New()
	if m.request.AppID == "" {
		v.Missing("app_id")
	}
	if m.request.Data == nil {
		v.Missing("data")
	}
	if _, err := json.Marshal(m.request); err != nil {
		v.Message("invalid format")
	}
	return v.Error()
}
