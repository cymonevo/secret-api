package main

import (
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

	handlers := provider.GetArticleHandler()
	log.Fatalf(mainTag, "Aborting...", http.ListenAndServe(":8000", handlers.Register()))
}
