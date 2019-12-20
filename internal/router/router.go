package router

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/cymon1997/go-backend/internal/log"
	"github.com/cymon1997/go-backend/internal/render"
	"github.com/gorilla/mux"
)

type Router interface {
	Handle(path string, method string, f func(ctx context.Context, r *http.Request) (RenderRequest, error))
	HandleJSON(path string, method string, f func(ctx context.Context, r *http.Request) (interface{}, error))
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type RenderRequest struct {
	Template string
	Data     interface{}
}

type handlerImpl struct {
	router       *mux.Router
	renderEngine render.Client
}

func NewRouter(renderEngine render.Client) Router {
	router := mux.NewRouter()
	return &handlerImpl{
		router:       router,
		renderEngine: renderEngine,
	}
}

func (h *handlerImpl) Handle(path string, method string, f func(ctx context.Context, r *http.Request) (RenderRequest, error)) {
	ctx := context.Background()
	handler := func(w http.ResponseWriter, r *http.Request) {
		request, _ := f(ctx, r)
		h.renderEngine.Render(w, request.Template, request.Data)
	}
	h.router.HandleFunc(path, handler).Methods(method)
}

func (h *handlerImpl) HandleJSON(path string, method string, f func(ctx context.Context, r *http.Request) (interface{}, error)) {
	ctx := context.Background()
	handler := func(w http.ResponseWriter, r *http.Request) {
		result, err := f(ctx, r)
		var response *Response
		if err != nil {
			response = h.setResponse(http.StatusInternalServerError, "internal server", nil)
		}
		response = h.setResponse(http.StatusOK, "success", result)
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			log.ErrorDetail(log.TagHandler, "error encode response", err)
		}
	}
	h.router.HandleFunc(path, handler).Methods(method)
}

func (h *handlerImpl) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

func (h *handlerImpl) setResponse(status int, message string, payload interface{}) *Response {
	return &Response{
		Status:  status,
		Message: message,
		Payload: payload,
	}
}
