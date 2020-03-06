package kloudless

import (
	"context"
	"errors"
	"fmt"

	"github.com/cymonevo/secret-api/handler"
	"github.com/cymonevo/secret-api/internal/log"
	"github.com/cymonevo/secret-api/sdk"
)

const getFolderContents = "/accounts/%d/storage/folders/%s/contents"

type GetFolderContentsRequest struct {
	FolderID string
}

type GetFolderContentsResponse struct {
	Count    int         `json:"count"`
	Type     string      `json:"type"`
	Objects  []File      `json:"objects"`
	Page     int         `json:"page"`
	NextPage interface{} `json:"next_page"`
}

func (c *clientImpl) GetFolderContents(ctx context.Context, req GetFolderContentsRequest) (GetFolderContentsResponse, error) {
	resp, err := c.client.Get(fmt.Sprintf(getFolderContents, c.AccID, req.FolderID), nil, c.headers(nil))
	if err != nil {
		log.ErrorDetail("GetFolderContents", "error get folder contents", err)
		return GetFolderContentsResponse{}, err
	}
	if !sdk.IsSuccess(resp.StatusCode) {
		log.Warnf("GetFolderContents", "status %d %s", resp.StatusCode, resp.Status)
		return GetFolderContentsResponse{}, errors.New(resp.Status)
	}
	var data GetFolderContentsResponse
	err = handler.ParseBody(resp.Body, &data)
	if err != nil {
		log.ErrorDetail("GetFolderContents", "error parse body", err)
		return GetFolderContentsResponse{}, err
	}
	log.Infof("GetFolderContents", "get folder contents success %+v", data)
	return data, nil
}
