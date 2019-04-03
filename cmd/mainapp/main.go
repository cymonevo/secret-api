package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/cymon1997/go-backend/internal/handler"
	article "github.com/cymon1997/go-backend/module/article/model"
)

var (
	articleFactory article.Factory
)

func main() {
	router := handler.NewRouter()
	router.Handle("/", http.MethodGet, Index)
	fmt.Println("Listen at port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func init() {
	articleFactory = article.NewFactory()
}

func Index(ctx context.Context, r *http.Request) (interface{}, error) {
	return struct {
		Version string `json:"version"`
		Build   string `json:"build_version"`
	}{
		Version: "0.0.1",
		Build:   "alpha",
	}, nil
}
