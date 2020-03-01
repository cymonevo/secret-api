package repo

import (
	"github.com/cymon1997/go-backend/internal/database"
	"github.com/jmoiron/sqlx"
)

type DBRepo interface {
	GetDB() *sqlx.DB
	SetDB(client database.Client)
}

type BaseDBRepo struct {
	dbClient database.Client
}

func (d *BaseDBRepo) SetDB(client database.Client) {
	d.dbClient = client
}

func (d *BaseDBRepo) GetDB() *sqlx.DB {
	return d.dbClient.GetInstance()
}
