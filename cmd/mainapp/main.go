package main

import (
	"net/http"

	"github.com/cymon1997/go-backend/internal/log"
	"github.com/cymon1997/go-backend/provider"
)

func main() {
	consumers := provider.GetArticleConsumers()
	for _, c := range consumers {
		c.Consume()
	}

	handlers := provider.GetArticleHandler()
	log.Fatalf("Aborting...", http.ListenAndServe(":8000", handlers.Register()))
}
