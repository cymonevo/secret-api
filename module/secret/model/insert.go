package model

import (
	"context"
	"errors"

	"github.com/cymonevo/secret-api/entity"
	"github.com/cymonevo/secret-api/internal/encoding/json"
	"github.com/cymonevo/secret-api/internal/log"
	"github.com/cymonevo/secret-api/internal/util"
	"github.com/cymonevo/secret-api/module/secret/repo"
)

const insertSecretTag = "Secret|Insert"

type InsertSecretModel struct {
	dbRepo  repo.DBRepo
	request entity.InsertSecretRequest
}

func (m *InsertSecretModel) Do(ctx context.Context) (entity.InsertSecretResponse, error) {
	var response entity.InsertSecretResponse
	if err := m.Validate(ctx); err != nil {
		log.ErrorDetail(insertSecretTag, "error validation: %v", err)
		response.Message = err.Error()
		return response, err
	}
	app, err := m.dbRepo.GetApp(ctx, m.request.AppID)
	if err != nil {
		log.ErrorDetail(insertSecretTag, "error get app data: %v", err)
		response.Message = err.Error()
		return response, err
	}
	last, err := m.dbRepo.GetLastSecret(ctx, m.request.AppID)
	if err != nil {
		log.ErrorDetail(insertSecretTag, "error get last data: %v", err)
		response.Message = err.Error()
		return response, err
	}
	old, err := util.Decrypt(last.Data, &app.Secret)
	if err != nil {
		log.ErrorDetail(insertSecretTag, "error decrypt old data: %v", err)
		response.Message = err.Error()
		return response, err
	}
	new, err := util.Decrypt(m.request.Data, &app.Secret)
	if err != nil {
		log.ErrorDetail(insertSecretTag, "error decrypt new data: %v", err)
		response.Message = err.Error()
		return response, err
	}
	result, err := json.Merge(old, new, json.OptionReplace)
	if err != nil {
		log.ErrorDetail(insertSecretTag, "error merge data: %v", err)
		response.Message = err.Error()
		return response, err
	}
	last.Data, err = util.Encrypt(result, &app.Secret)
	if err != nil {
		log.ErrorDetail(insertSecretTag, "error encrypt result: %v", err)
		response.Message = err.Error()
		return response, err
	}
	err = m.dbRepo.InsertSecret(ctx, last)
	if err != nil {
		log.ErrorDetail(insertSecretTag, "error save to db: %v", err)
		response.Message = err.Error()
		return response, err
	}
	return response, nil
}

func (m *InsertSecretModel) Validate(ctx context.Context) error {
	if m.request.AppID == "" {
		return errors.New("invalid request")
	}
	if m.request.Data == nil {
		return errors.New("invalid request")
	}
}
