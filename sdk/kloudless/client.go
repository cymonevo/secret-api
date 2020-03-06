package kloudless

import (
	"context"
	"fmt"

	"github.com/cymonevo/secret-api/internal/util"
	"github.com/cymonevo/secret-api/sdk"
)

type Client interface {
	GetAccounts(ctx context.Context) (interface{}, error)
	CreateFolder(ctx context.Context, req CreateFolderRequest) (CreateFolderResponse, error)
	GetFolderContents(ctx context.Context, req GetFolderContentsRequest) (GetFolderContentsResponse, error)
	GetFileInfo(ctx context.Context, req GetFileInfoRequest) (GetFileInfoResponse, error)
	DownloadFile(ctx context.Context, req DownloadFileRequest) (DownloadFileResponse, error)
	UploadFile(ctx context.Context, req UploadFileRequest) (UploadFileResponse, error)
}

type Config struct {
	sdk.Config
	APIKey string
	AccID  int
}

type clientImpl struct {
	client sdk.Client
	APIKey string
	AccID  int
}

func New(cfg Config) *clientImpl {
	return &clientImpl{
		client: sdk.New(sdk.Config{
			Timeout: cfg.Timeout,
			URL:     cfg.URL,
		}),
		APIKey: cfg.APIKey,
		AccID:  cfg.AccID,
	}
}

func (c *clientImpl) headers(headers map[string]string) map[string]string {
	return util.CombineMapString(map[string]string{
		"Authorization": fmt.Sprint("APIKey ", c.APIKey),
	}, headers)
}
