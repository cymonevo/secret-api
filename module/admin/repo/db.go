package repo

import (
	"context"

	"github.com/cymonevo/secret-api/entity"
	"github.com/cymonevo/secret-api/internal/base/repo"
)

type DBRepo interface {
	InsertApp(ctx context.Context, data entity.AppData) error
	GetApp(ctx context.Context, id string) (entity.AppData, error)
}

type AdminDBRepo struct {
	db *repo.BaseDBRepo
}

func NewAdminDBRepo(db *repo.BaseDBRepo) *AdminDBRepo {
	return &AdminDBRepo{
		db: db,
	}
}

func (r *AdminDBRepo) InsertApp(ctx context.Context, data entity.AppData) error {
	_, err := r.db.GetDB().NamedExecContext(ctx, insertAppQuery, data)
	if err != nil {
		return err
	}
	return nil
}

func (r *AdminDBRepo) GetApp(ctx context.Context, id string) (entity.AppData, error) {
	var result entity.AppData
	err := r.db.GetDB().GetContext(ctx, &result, getAppQuery, id)
	if err != nil {
		return result, err
	}
	return result, nil
}
