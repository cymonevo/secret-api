package handler

import (
	"context"
	"net/http"

	"github.com/cymonevo/secret-api/entity"
	"github.com/cymonevo/secret-api/internal/log"
	"github.com/cymonevo/secret-api/internal/router"
	"github.com/cymonevo/secret-api/module/secret/model"
)

const secretHandlerTag = "Handler|Secret"

type SecretHandler struct {
	factory *model.Factory
}

func NewSecretHandler(factory *model.Factory) *SecretHandler {
	return &SecretHandler{
		factory: factory,
	}
}

func (h *SecretHandler) Register(router router.Router) {
	router.HandleJSON("/secret", http.MethodPut, h.insertSecret)
}

func (h *SecretHandler) insertSecret(ctx context.Context, r *http.Request) (interface{}, error) {
	var data entity.InsertSecretRequest
	err := ParseBody(r.Body, &data)
	if err != nil {
		log.ErrorDetail(secretHandlerTag, "error parse body: %v", err)
		return nil, err
	}
	return h.factory.NewInsertSecretModel(data).Do(ctx)
}
