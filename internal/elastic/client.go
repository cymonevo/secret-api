package elastic

import (
	"context"
	"net/http"

	"github.com/cymon1997/go-backend/internal/config"
	"github.com/cymon1997/go-backend/internal/log"
	esconfig "github.com/olivere/elastic/config"
	"gopkg.in/olivere/elastic.v6"
)

type ESClient interface {
	Dial(ctx context.Context) (*elastic.PingResult, error)
	IndexExists(ctx context.Context, index string) (bool, error)
	CreateIndex(ctx context.Context, index string) (*elastic.IndicesCreateResult, error)
	GetByID(ctx context.Context, index string, dataType string, id string) (*elastic.GetResult, error)
	PutJSON(ctx context.Context, index string, dataType string, data interface{}) (*elastic.IndexResponse, error)
	PutJSONWithID(ctx context.Context, index string, dataType string, id string, data interface{}) (*elastic.IndexResponse, error)
	PutString(ctx context.Context, index string, dataType string, data string) (*elastic.IndexResponse, error)
	PutStringWithID(ctx context.Context, index string, dataType string, id string, data string) (*elastic.IndexResponse, error)
	Update(ctx context.Context, index string, dataType string, id string, data interface{}, script *elastic.Script) (*elastic.UpdateResponse, error)
	DeleteIndex(ctx context.Context, index string) (*elastic.IndicesDeleteResponse, error)
	DeleteByID(ctx context.Context, index string, dataType string, id string) (*elastic.DeleteResponse, error)
	GetSearchInstance() *elastic.SearchService
	Flush(ctx context.Context, index string) (*elastic.IndicesFlushResponse, error)
}

type esClient struct {
	client *elastic.Client
	host   string
}

func NewESClient(cfg *config.ESConfig) ESClient {
	client, err := elastic.NewClientFromConfig(parseConfig(cfg))
	if err != nil {
		log.FatalDetail(log.TagES, "error create es client", err)
	}
	return &esClient{
		client: client,
		host:   parseAddress(cfg),
	}
}

func parseConfig(cfg *config.ESConfig) *esconfig.Config {
	return &esconfig.Config{
		URL: parseAddress(cfg),
	}
}

func (c *esClient) Dial(ctx context.Context) (*elastic.PingResult, error) {
	result, status, err := c.client.Ping(c.host).Do(ctx)
	if err != nil && status != http.StatusOK {
		log.ErrorDetail(log.TagES, "error dial es", err)
		return nil, err
	}
	return result, nil
}

func (c *esClient) IndexExists(ctx context.Context, index string) (bool, error) {
	exist, err := c.client.IndexExists(index).Do(ctx)
	if err != nil {
		log.ErrorDetail(log.TagES, "error check index", err)
		return false, err
	}
	return exist, nil
}

func (c *esClient) CreateIndex(ctx context.Context, index string) (*elastic.IndicesCreateResult, error) {
	result, err := c.client.CreateIndex(index).Do(ctx)
	if err != nil {
		log.ErrorDetail(log.TagES, "error create index", err)
		return nil, err
	}
	return result, nil
}

func (c *esClient) GetByID(ctx context.Context, index string, dataType string, id string) (*elastic.GetResult, error) {
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

func (c *esClient) PutJSON(ctx context.Context, index string, dataType string, data interface{}) (*elastic.IndexResponse, error) {
	result, err := c.client.Index().Index(index).Type(dataType).BodyJson(data).Do(ctx)
	if err != nil {
		log.ErrorDetail(log.TagES, "error put data to es", err)
		return nil, err
	}
	return result, nil
}

func (c *esClient) PutJSONWithID(ctx context.Context, index string, dataType string, id string, data interface{}) (*elastic.IndexResponse, error) {
	result, err := c.client.Index().Index(index).Type(dataType).Id(id).BodyJson(data).Do(ctx)
	if err != nil {
		log.ErrorDetail(log.TagES, "error put data to es", err)
		return nil, err
	}
	return result, nil
}

func (c *esClient) PutString(ctx context.Context, index string, dataType string, data string) (*elastic.IndexResponse, error) {
	result, err := c.client.Index().Index(index).Type(dataType).BodyString(data).Do(ctx)
	if err != nil {
		log.ErrorDetail(log.TagES, "error put data to es", err)
		return nil, err
	}
	return result, nil
}

func (c *esClient) PutStringWithID(ctx context.Context, index string, dataType string, id string, data string) (*elastic.IndexResponse, error) {
	result, err := c.client.Index().Index(index).Type(dataType).Id(id).BodyString(data).Do(ctx)
	if err != nil {
		log.ErrorDetail(log.TagES, "error put data to es", err)
		return nil, err
	}
	return result, nil
}

func (c *esClient) Update(ctx context.Context, index string, dataType string, id string, data interface{}, script *elastic.Script) (*elastic.UpdateResponse, error) {
	result, err := c.client.Update().Index(index).Type(dataType).Id(id).Script(script).Upsert(data).Do(ctx)
	if err != nil {
		log.ErrorDetail(log.TagES, "error update data to es", err)
		return nil, err
	}
	return result, nil
}

func (c *esClient) DeleteIndex(ctx context.Context, index string) (*elastic.IndicesDeleteResponse, error) {
	result, err := c.client.DeleteIndex(index).Do(ctx)
	if err != nil {
		log.ErrorDetail(log.TagES, "error delete index", err)
		return nil, err
	}
	return result, nil
}

func (c *esClient) DeleteByID(ctx context.Context, index string, dataType string, id string) (*elastic.DeleteResponse, error) {
	result, err := c.client.Delete().Index(index).Type(dataType).Id(id).Do(ctx)
	if err != nil {
		log.ErrorDetail(log.TagES, "error delete data by id", err)
		return nil, err
	}
	return result, nil
}

func (c *esClient) GetSearchInstance() *elastic.SearchService {
	return c.client.Search()
}

func (c *esClient) Flush(ctx context.Context, index string) (*elastic.IndicesFlushResponse, error) {
	result, err := c.client.Flush().Index(index).Do(ctx)
	if err != nil {
		log.ErrorDetail(log.TagES, "error flush index to es", err)
		return nil, err
	}
	return result, nil
}
