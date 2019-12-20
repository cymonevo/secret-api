package elastic

import (
	"context"
	"net/http"

	"github.com/cymon1997/go-backend/internal/config"
	"github.com/cymon1997/go-backend/internal/log"
	esconfig "github.com/olivere/elastic/config"
	"gopkg.in/olivere/elastic.v6"
)

type Client interface {
	Dial(ctx context.Context) (*elastic.PingResult, error)
	IndexExists(ctx context.Context, index string) (bool, error)
	CreateIndex(ctx context.Context, index string) (*elastic.IndicesCreateResult, error)
	GetByID(ctx context.Context, index string, dataType string, id string) (*elastic.GetResult, error)
	PutJSON(ctx context.Context, index string, dataType string, data interface{}) (*elastic.IndexResponse, error)
	PutJSONWithID(ctx context.Context, index string, dataType string, id string, data interface{}) (*elastic.IndexResponse, error)
	PutString(ctx context.Context, index string, dataType string, data string) (*elastic.IndexResponse, error)
	PutStringWithID(ctx context.Context, index string, dataType string, id string, data string) (*elastic.IndexResponse, error)
	Update(ctx context.Context, index string, dataType string, id string, data interface{}) (*elastic.UpdateResponse, error)
	UpdateWithScript(ctx context.Context, index string, dataType string, id string, data interface{}, script *elastic.Script) (*elastic.UpdateResponse, error)
	DeleteIndex(ctx context.Context, index string) (*elastic.IndicesDeleteResponse, error)
	DeleteByID(ctx context.Context, index string, dataType string, id string) (*elastic.DeleteResponse, error)
	GetSearchInstance() *elastic.SearchService
	Flush(ctx context.Context, index string) (*elastic.IndicesFlushResponse, error)
}

type clientImpl struct {
	client *elastic.Client
	host   string
}

func NewESClient(cfg config.ESConfig) *clientImpl {
	client, err := elastic.NewClientFromConfig(parseConfig(cfg))
	if err != nil {
		log.FatalDetail(log.TagES, "error create es client", err)
	}
	return &clientImpl{
		client: client,
		host:   parseAddress(cfg),
	}
}

func parseConfig(cfg config.ESConfig) *esconfig.Config {
	return &esconfig.Config{
		URL: parseAddress(cfg),
	}
}

func (c *clientImpl) Dial(ctx context.Context) (*elastic.PingResult, error) {
	result, status, err := c.client.Ping(c.host).Do(ctx)
	if err != nil && status != http.StatusOK {
		log.ErrorDetail(log.TagES, "error dial es", err)
		return nil, err
	}
	return result, nil
}

func (c *clientImpl) IndexExists(ctx context.Context, index string) (bool, error) {
	exist, err := c.client.IndexExists(index).Do(ctx)
	if err != nil {
		log.ErrorDetail(log.TagES, "error check index", err)
		return false, err
	}
	return exist, nil
}

func (c *clientImpl) CreateIndex(ctx context.Context, index string) (*elastic.IndicesCreateResult, error) {
	result, err := c.client.CreateIndex(index).Do(ctx)
	if err != nil {
		log.ErrorDetail(log.TagES, "error create index", err)
		return nil, err
	}
	return result, nil
}

func (c *clientImpl) GetByID(ctx context.Context, index string, dataType string, id string) (*elastic.GetResult, error) {
	result, err := c.client.Get().Index(index).Type(dataType).Id(id).Do(ctx)
	if err != nil {
		log.ErrorDetail(log.TagES, "error get id", err)
		return nil, err
	}
	if !result.Found {
		log.WarnDetail(log.TagES, "error id not found", err)
		return nil, nil
	}
	return result, nil
}

func (c *clientImpl) PutJSON(ctx context.Context, index string, dataType string, data interface{}) (*elastic.IndexResponse, error) {
	result, err := c.client.Index().Index(index).Type(dataType).BodyJson(data).Do(ctx)
	if err != nil {
		log.ErrorDetail(log.TagES, "error put data to es", err)
		return nil, err
	}
	return result, nil
}

func (c *clientImpl) PutJSONWithID(ctx context.Context, index string, dataType string, id string, data interface{}) (*elastic.IndexResponse, error) {
	result, err := c.client.Index().Index(index).Type(dataType).Id(id).BodyJson(data).Do(ctx)
	if err != nil {
		log.ErrorDetail(log.TagES, "error put data to es", err)
		return nil, err
	}
	return result, nil
}

func (c *clientImpl) PutString(ctx context.Context, index string, dataType string, data string) (*elastic.IndexResponse, error) {
	result, err := c.client.Index().Index(index).Type(dataType).BodyString(data).Do(ctx)
	if err != nil {
		log.ErrorDetail(log.TagES, "error put data to es", err)
		return nil, err
	}
	return result, nil
}

func (c *clientImpl) PutStringWithID(ctx context.Context, index string, dataType string, id string, data string) (*elastic.IndexResponse, error) {
	result, err := c.client.Index().Index(index).Type(dataType).Id(id).BodyString(data).Do(ctx)
	if err != nil {
		log.ErrorDetail(log.TagES, "error put data to es", err)
		return nil, err
	}
	return result, nil
}

func (c *clientImpl) Update(ctx context.Context, index string, dataType string, id string, data interface{}) (*elastic.UpdateResponse, error) {
	result, err := c.client.Update().Index(index).Type(dataType).Id(id).Doc(data).DocAsUpsert(true).Do(ctx)
	if err != nil {
		log.ErrorDetail(log.TagES, "error update data to es", err)
		return nil, err
	}
	return result, nil
}

func (c *clientImpl) UpdateWithScript(ctx context.Context, index string, dataType string, id string, data interface{}, script *elastic.Script) (*elastic.UpdateResponse, error) {
	result, err := c.client.Update().Index(index).Type(dataType).Id(id).Script(script).Upsert(data).Do(ctx)
	if err != nil {
		log.ErrorDetail(log.TagES, "error update data to es", err)
		return nil, err
	}
	return result, nil
}

func (c *clientImpl) DeleteIndex(ctx context.Context, index string) (*elastic.IndicesDeleteResponse, error) {
	result, err := c.client.DeleteIndex(index).Do(ctx)
	if err != nil {
		log.ErrorDetail(log.TagES, "error delete index", err)
		return nil, err
	}
	return result, nil
}

func (c *clientImpl) DeleteByID(ctx context.Context, index string, dataType string, id string) (*elastic.DeleteResponse, error) {
	result, err := c.client.Delete().Index(index).Type(dataType).Id(id).Do(ctx)
	if err != nil {
		log.ErrorDetail(log.TagES, "error delete data by id", err)
		return nil, err
	}
	return result, nil
}

func (c *clientImpl) GetSearchInstance() *elastic.SearchService {
	return c.client.Search()
}

func (c *clientImpl) Flush(ctx context.Context, index string) (*elastic.IndicesFlushResponse, error) {
	result, err := c.client.Flush().Index(index).Do(ctx)
	if err != nil {
		log.ErrorDetail(log.TagES, "error flush index to es", err)
		return nil, err
	}
	return result, nil
}
