package repo

import (
	"github.com/cymonevo/secret-api/internal/database"
	"github.com/jmoiron/sqlx"
)

type BaseDBRepo struct {
	dbClient database.Client
}

func NewBaseDBRepo(db database.Client) *BaseDBRepo {
	return &BaseDBRepo{
		dbClient: db,
	}
}

func (d *BaseDBRepo) GetDB() *sqlx.DB {
	return d.dbClient.GetInstance()
}

//TODO: Implement basic DB operation

//TODO: Implement statement management
