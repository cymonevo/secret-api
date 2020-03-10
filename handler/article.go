package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cymonevo/secret-api/internal/log"
	"github.com/cymonevo/secret-api/internal/router"
	"github.com/cymonevo/secret-api/module/article/model"
)

type articleHandlerImpl struct {
	factory model.Factory
}

func NewArticleHandler(factory model.Factory) *articleHandlerImpl {
	return &articleHandlerImpl{
		factory: factory,
	}
}

func (h *articleHandlerImpl) Register(router router.Router) {
	router.HandleJSON("/article", http.MethodGet, h.index)
	router.HandleView("/article/view", http.MethodGet, h.view)
	//test endpoints
	router.HandleJSON("/article/get", http.MethodGet, h.get)
	router.HandleJSON("/article/post", http.MethodPost, h.post)
}

func (h *articleHandlerImpl) get(ctx context.Context, r *http.Request) (interface{}, error) {
	auth := r.Header.Get("Authorization")
	query := GetQueryParam(r, "data")
	return fmt.Sprint("GET", "\nAUTH: ", auth, "\nDATA: ", query), nil
}

func (h *articleHandlerImpl) post(ctx context.Context, r *http.Request) (interface{}, error) {
	auth := r.Header.Get("Authorization")
	var data interface{}
	err := ParseBody(r.Body, &data)
	if err != nil {
		log.ErrorDetail("Article", "error parse request body", err)
		return nil, err
	}
	return fmt.Sprint("POST", "\nAUTH: ", auth, "\nDATA: ", data), nil
}

func (h *articleHandlerImpl) index(ctx context.Context, r *http.Request) (interface{}, error) {
	return struct {
		Version string `json:"version"`
		Build   string `json:"build_version"`
	}{
		Version: "0.0.1",
		Build:   "alpha",
	}, nil
}

func (h *articleHandlerImpl) view(ctx context.Context, r *http.Request) (router.RenderRequest, error) {
	type invoice struct {
		Invoice string
		OrderID string
	}
	return router.RenderRequest{
		Template: "invoice.html",
		Data: invoice{
			Invoice: "INV/2018/123",
			OrderID: "123",
		},
	}, nil
}

func (h *articleHandlerImpl) health(ctx context.Context, r *http.Request) (interface{}, error) {
	return h.factory.NewHealthModel().Do(ctx)
}
