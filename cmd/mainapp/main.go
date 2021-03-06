package main

import (
	"fmt"
	"net/http"

	"github.com/cymonevo/secret-api/internal/log"
	"github.com/cymonevo/secret-api/provider"
)

const mainTag = "Main"

func main() {
	//consumers := provider.GetArticleConsumers()
	//for _, c := range consumers {
	//	c.Consume()
	//}
	cfg := provider.GetAppConfig()
	log.Infof(mainTag, "app is running at :%d", cfg.AppConfig.Port)
	log.Fatalf(mainTag, "aborting...", http.ListenAndServe(fmt.Sprintf("%s:%d",
		cfg.AppConfig.Host, cfg.AppConfig.Port), provider.SetupHandler()))
}
