package repo

import (
	"context"
	"github.com/cymonevo/secret-api/entity"

	"github.com/cymonevo/secret-api/internal/base/repo"
)

type DBRepo interface {
	InsertApp(ctx context.Context, data entity.AppData) error
	GetApp(ctx context.Context, id string) (entity.AppData, error)
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

func (r *SecretDBRepo) InsertApp(ctx context.Context, data entity.AppData) error {
	_, err := r.db.GetDB().NamedExecContext(ctx, insertAppQuery, data)
	if err != nil {
		return err
	}
	return nil
}

func (r *SecretDBRepo) GetApp(ctx context.Context, id string) (entity.AppData, error) {
	var result entity.AppData
	err := r.db.GetDB().GetContext(ctx, &result, getAppQuery, id)
	if err != nil {
		return result, err
	}
	return result, nil
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
		return result, err
	}
	return result, nil
}

func (r *SecretDBRepo) GetLastSecret(ctx context.Context, appID string) (entity.SecretData, error) {
	var result entity.SecretData
	err := r.db.GetDB().GetContext(ctx, &result, getLastSecretQuery, appID)
	if err != nil {
		return result, err
	}
	return result, nil
}
