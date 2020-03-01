package kloudless

import (
	"context"
	"errors"
	"fmt"

	"github.com/cymon1997/go-backend/handler"
	"github.com/cymon1997/go-backend/internal/log"
	"github.com/cymon1997/go-backend/sdk"
)

const downloadFile = "/accounts/%d/storage/files/%s/contents"

type DownloadFileRequest struct {
	FileID string
}

type DownloadFileResponse struct {
	RawFile []byte `json:"raw"`
}

func (c *clientImpl) DownloadFile(ctx context.Context, req DownloadFileRequest) (DownloadFileResponse, error) {
	resp, err := c.client.Get(fmt.Sprintf(downloadFile, c.AccID, req.FileID), nil, c.headers(nil))
	if err != nil {
		log.ErrorDetail("DownloadFile", "error download file", err)
		return DownloadFileResponse{}, err
	}
	if !sdk.IsSuccess(resp.StatusCode) {
		log.Warnf("DownloadFile", "status %d %s", resp.StatusCode, resp.Status)
		return DownloadFileResponse{}, errors.New(resp.Status)
	}
	var data DownloadFileResponse
	data.RawFile, err = handler.ParseFile(resp.Body)
	if err != nil {
		log.ErrorDetail("DownloadFile", "error parse file", err)
		return DownloadFileResponse{}, err
	}
	log.Infof("DownloadFile", "download file success %+v", data)
	return data, nil
}
