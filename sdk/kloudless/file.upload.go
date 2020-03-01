package kloudless

import (
	"context"
	"errors"
	"fmt"

	"github.com/cymon1997/go-backend/handler"
	"github.com/cymon1997/go-backend/internal/log"
	"github.com/cymon1997/go-backend/sdk"
)

const uploadFile = "/accounts/%d/storage/files"

type UploadFileRequest struct {
	DestID  string `json:"parent_id"`
	Name    string `json:"name"`
	RawFile []byte `json:"raw"`
}

type UploadFileResponse struct {
	File
}

func (c *clientImpl) UploadFile(ctx context.Context, req UploadFileRequest) (UploadFileResponse, error) {
	resp, err := c.client.PostRaw(fmt.Sprintf(uploadFile, c.AccID), req.RawFile, c.headers(map[string]string{
		"X-Kloudless-Metadata": fmt.Sprintf(`{"name": "%s", "parent_id": "%s"}`, req.Name, req.DestID),
	}))
	if err != nil {
		log.ErrorDetail("UploadFile", "error upload file", err)
		return UploadFileResponse{}, err
	}
	if !sdk.IsSuccess(resp.StatusCode) {
		log.Warnf("UploadFile", "status %d %s", resp.StatusCode, resp.Status)
		return UploadFileResponse{}, errors.New(resp.Status)
	}
	var data UploadFileResponse
	err = handler.ParseBody(resp.Body, &data)
	if err != nil {
		log.ErrorDetail("UploadFile", "error parse body", err)
		return UploadFileResponse{}, err
	}
	log.Infof("UploadFile", "upload file success %+v", data)
	return data, nil
}
