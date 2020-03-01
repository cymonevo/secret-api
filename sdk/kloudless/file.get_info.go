package kloudless

import (
	"context"
	"errors"
	"fmt"

	"github.com/cymon1997/go-backend/handler"
	"github.com/cymon1997/go-backend/internal/log"
	"github.com/cymon1997/go-backend/sdk"
)

const getFileInfo = "/accounts/%d/storage/files/%s"

type GetFileInfoRequest struct {
	FileID string
}

type GetFileInfoResponse struct {
	File
}

func (c *clientImpl) GetFileInfo(ctx context.Context, req GetFileInfoRequest) (GetFileInfoResponse, error) {
	resp, err := c.client.Get(fmt.Sprintf(getFileInfo, c.AccID, req.FileID), nil, c.headers(nil))
	if err != nil {
		log.ErrorDetail("GetFileInfo", "error get file info", err)
		return GetFileInfoResponse{}, err
	}
	if !sdk.IsSuccess(resp.StatusCode) {
		log.Warnf("GetFileInfo", "status %d %s", resp.StatusCode, resp.Status)
		return GetFileInfoResponse{}, errors.New(resp.Status)
	}
	var data GetFileInfoResponse
	err = handler.ParseBody(resp.Body, &data)
	if err != nil {
		log.ErrorDetail("GetFileInfo", "error parse body", err)
		return GetFileInfoResponse{}, err
	}
	log.Infof("GetFileInfo", "get file info success %+v", data)
	return data, nil
}
