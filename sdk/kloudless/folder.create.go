package kloudless

import (
	"context"
	"errors"
	"fmt"

	"github.com/cymonevo/secret-api/handler"
	"github.com/cymonevo/secret-api/internal/log"
	"github.com/cymonevo/secret-api/sdk"
)

const createFolder = "/accounts/%d/storage/folders"

type CreateFolderRequest struct {
	DestID string `json:"parent_id"`
	Name   string `json:"name"`
}

type CreateFolderResponse struct {
	File
}

func (c *clientImpl) CreateFolder(ctx context.Context, req CreateFolderRequest) (CreateFolderResponse, error) {
	resp, err := c.client.Post(fmt.Sprintf(createFolder, c.AccID), req, c.headers(nil))
	if err != nil {
		log.ErrorDetail("CreateFolder", "error create folder", err)
		return CreateFolderResponse{}, err
	}
	if !sdk.IsSuccess(resp.StatusCode) {
		log.Warnf("CreateFolder", "status %d %s", resp.StatusCode, resp.Status)
		return CreateFolderResponse{}, errors.New(resp.Status)
	}
	var data CreateFolderResponse
	err = handler.ParseBody(resp.Body, &data)
	if err != nil {
		log.ErrorDetail("CreateFolder", "error parse body", err)
		return CreateFolderResponse{}, err
	}
	log.Infof("CreateFolder", "create folder success %+v", data)
	return data, nil
}
