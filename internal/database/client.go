package database

import (
	"github.com/cymon1997/go-backend/internal/config"
	"github.com/cymon1997/go-backend/internal/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Client interface {
	Dial() error
	GetInstance() *sqlx.DB
}

type clientImpl struct {
	db *sqlx.DB
}

func New(cfg config.DBConfig) *clientImpl {
	db, err := sqlx.Open(cfg.Driver, parseDSN(cfg))
	if err != nil {
		log.FatalDetail(log.TagDB, "error create db clientImpl", err)
	}
	return &clientImpl{
		db: db,
	}
}

func (c *clientImpl) Dial() error {
	err := c.db.Ping()
	if err != nil {
		log.ErrorDetail(log.TagDB, "error ping db", err)
		return err
	}
	return nil
}

func (c *clientImpl) GetInstance() *sqlx.DB {
	return c.db
}
