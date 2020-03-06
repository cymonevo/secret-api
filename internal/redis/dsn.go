package redis

import (
	"fmt"

	"github.com/cymonevo/secret-api/internal/config"
)

func parseAddress(cfg *config.RedisConfig) string {
	//ADDR FORMAT : 127.0.0.1:3306
	return fmt.Sprintf("%s:%v", cfg.Address, cfg.Port)
}
