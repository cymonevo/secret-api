package model

import (
	"context"
	"database/sql"
	"errors"

	"github.com/cymonevo/secret-api/entity"
	"github.com/cymonevo/secret-api/internal/log"
	"github.com/cymonevo/secret-api/internal/util"
	"github.com/cymonevo/secret-api/module/admin/repo"
)

const registerTag = "Admin|Register"

type RegisterModel struct {
	dbRepo  repo.DBRepo
	request entity.RegisterRequest
}

func (m *RegisterModel) Do(ctx context.Context) (entity.RegisterResponse, error) {
	var response entity.RegisterResponse
	if err := m.Validate(ctx); err != nil {
		log.ErrorDetail(registerTag, "error validation: %v", err)
		response.Message = err.Error()
		return response, err
	}
	_, err := m.dbRepo.GetApp(ctx, m.request.AppID)
	if err != nil && err != sql.ErrNoRows {
		log.ErrorDetail(registerTag, "error get app data: %v", err)
		response.Message = err.Error()
		return response, err
	}
	if err == nil {
		log.WarnDetail(registerTag, "app_id already exists")
		response.Message = "app already registered"
		return response, err
	}
	key := util.NewEncryptionKey()[:]
	err = m.dbRepo.InsertApp(ctx, entity.AppData{
		AppID:  m.request.AppID,
		Secret: key,
	})
	if err != nil {
		log.ErrorDetail(registerTag, "error save to db: %v", err)
		response.Message = err.Error()
		return response, err
	}
	response.Secret = key
	return response, nil
}

func (m *RegisterModel) Validate(ctx context.Context) error {
	if m.request.AppID == "" {
		return errors.New("invalid request")
	}
	return nil
}
