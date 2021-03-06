package repo

import (
	"context"
	"database/sql"

	"github.com/cymonevo/secret-api/entity"
	"github.com/cymonevo/secret-api/internal/base/repo"
	"github.com/cymonevo/secret-api/internal/errors"
)

type DBRepo interface {
	InsertSecret(ctx context.Context, data entity.SecretData) error
	GetAllSecret(ctx context.Context, appID string, limit int) ([]entity.SecretData, error)
	GetLastSecret(ctx context.Context, appID string) (entity.SecretData, error)
}

type SecretDBRepo struct {
	db *repo.BaseDBRepo
}

func NewSecretDBRepo(db *repo.BaseDBRepo) *SecretDBRepo {
	return &SecretDBRepo{
		db: db,
	}
}

func (r *SecretDBRepo) InsertSecret(ctx context.Context, data entity.SecretData) error {
	_, err := r.db.GetDB().NamedExecContext(ctx, insertSecretQuery, data)
	if err != nil {
		return err
	}
	return nil
}

func (r *SecretDBRepo) GetAllSecret(ctx context.Context, appID string, limit int) ([]entity.SecretData, error) {
	var result []entity.SecretData
	err := r.db.GetDB().SelectContext(ctx, &result, getAllSecretQuery, appID, limit)
	if err != nil {
		if err == sql.ErrNoRows {
			return result, errors.New(errors.NoDataFound)
		}
		return result, err
	}
	return result, nil
}

func (r *SecretDBRepo) GetLastSecret(ctx context.Context, appID string) (entity.SecretData, error) {
	var result entity.SecretData
	err := r.db.GetDB().GetContext(ctx, &result, getLastSecretQuery, appID)
	if err != nil {
		if err == sql.ErrNoRows {
			return result, errors.New(errors.NoDataFound)
		}
		return result, err
	}
	return result, nil
}
