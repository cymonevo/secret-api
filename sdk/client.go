package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/cymon1997/go-backend/handler"
	"github.com/cymon1997/go-backend/internal/log"
	"github.com/google/go-querystring/query"
)

type Client interface {
	Get(url string, params interface{}, headers map[string]string) (*http.Response, error)
	Post(url string, body interface{}, headers map[string]string) (*http.Response, error)
	PostRaw(uri string, body []byte, headers map[string]string) (*http.Response, error)
}

type clientImpl struct {
	client  http.Client
	baseURL string
}

func New(cfg Config) *clientImpl {
	return &clientImpl{
		client: http.Client{
			Timeout: time.Duration(cfg.Timeout) * time.Second,
		},
		baseURL: cfg.URL,
	}
}

func (c *clientImpl) Get(uri string, params interface{}, headers map[string]string) (*http.Response, error) {
	v, _ := query.Values(params)
	return c.execute(http.MethodGet, fmt.Sprintf("%s%s?%s", c.baseURL, uri, v), nil, headers)
}

func (c *clientImpl) Post(uri string, body interface{}, headers map[string]string) (*http.Response, error) {
	raw, err := json.Marshal(body)
	if err != nil {
		log.ErrorDetail("SDK", "error parse request data", err)
		return nil, err
	}
	return c.execute(http.MethodPost, fmt.Sprint(c.baseURL, uri), raw, headers)
}

func (c *clientImpl) PostRaw(uri string, body []byte, headers map[string]string) (*http.Response, error) {
	return c.execute(http.MethodPost, fmt.Sprint(c.baseURL, uri), body, headers)
}

func (c *clientImpl) execute(method, url string, body []byte, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		log.ErrorDetail("SDK", "error create get request", err)
		return nil, err
	}
	handler.SetHeaders(req, headers)
	return c.client.Do(req)
}
