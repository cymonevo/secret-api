package handler

import (
	"context"
	"net/http"

	"github.com/cymonevo/secret-api/entity"
	"github.com/cymonevo/secret-api/internal/log"
	"github.com/cymonevo/secret-api/internal/router"
	"github.com/cymonevo/secret-api/module/admin/model"
)

const adminHandlerTag = "Handler|Admin"

type AdminHandler struct {
	factory *model.Factory
}

func NewAdminHandler(factory *model.Factory) *AdminHandler {
	return &AdminHandler{
		factory: factory,
	}
}

func (h *AdminHandler) Register(router router.Router) {
	router.HandleJSON("/admin/register", http.MethodPost, h.register)
}

func (h *AdminHandler) register(ctx context.Context, r *http.Request) (interface{}, error) {
	var data entity.RegisterRequest
	err := ParseBody(r.Body, &data)
	if err != nil {
		log.ErrorDetail(adminHandlerTag, "error parse body: %v", err)
		return nil, err
	}
	return h.factory.NewRegisterModel(data).Do(ctx)
}
