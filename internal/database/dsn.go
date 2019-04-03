package database

import (
	"fmt"

	"github.com/cymon1997/go-backend/internal/config"
)

func parseDSN(cfg *config.DBConfig) string {
	//DSN FORMAT : mysql://root:root_pass@127.0.0.1:3306/test_db
	//LIB FORMAT : root:root_pass@127.0.0.1:3306/test_db
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
}
