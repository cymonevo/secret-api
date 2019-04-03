package redis

import (
	"fmt"

	"github.com/cymon1997/go-backend/internal/config"
)

func parseAddress(cfg *config.RedisConfig) string {
	//ADDR FORMAT : 127.0.0.1:3306
	return fmt.Sprintf("%s:%v", cfg.Address, cfg.Port)
}
