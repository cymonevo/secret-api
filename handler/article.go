package handler

import (
	"context"
	"github.com/cymon1997/go-backend/module/article/model"
	"net/http"

	"github.com/cymon1997/go-backend/internal/router"
)

type articleHandlerImpl struct {
	router  router.Router
	factory model.Factory
}

func NewArticleHandler(router router.Router, factory model.Factory) *articleHandlerImpl {
	return &articleHandlerImpl{
		router:  router,
		factory: factory,
	}
}

func (h *articleHandlerImpl) Register() router.Router {
	h.router.SetPrefix("/article")
	h.router.HandleJSON("", http.MethodGet, h.index)
	h.router.HandleView("/view", http.MethodGet, h.view)
	return h.router
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
	resp, err := h.factory.NewHealthModel().Call(ctx)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
