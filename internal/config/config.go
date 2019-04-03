package config

import (
	"fmt"
	"os"

	"github.com/cymon1997/go-backend/internal/log"
	"gopkg.in/gcfg.v1"
)

func LoadMainConfig() *MainConfig {
	var cfg MainConfig
	err := gcfg.ReadFileInto(&cfg, fmt.Sprintf(configLocation, detectEnv()))
	if err != nil {
		log.FatalDetail(log.TagConfig, "error load config", err)
	}
	return &cfg
}

func detectEnv() string {
	switch os.Getenv(env) {
	case staging:
		return staging
	case production:
		return production
	default:
		return local
	}
}
