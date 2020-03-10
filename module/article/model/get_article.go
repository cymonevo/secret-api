package model

import (
	"context"
	"time"

	"github.com/cymonevo/secret-api/entity"
	"github.com/cymonevo/secret-api/internal/log"
	"gopkg.in/guregu/null.v3"
)

const getArticleTag = "Article|Get"

type GetArticleModel struct {
}

func (m *GetArticleModel) Do(ctx context.Context) (entity.GetArticleResponse, error) {
	var response entity.GetArticleResponse
	err := m.Validate(ctx)
	if err != nil {
		log.ErrorDetail(getArticleTag, "error validation: %v", err)
		response.Message = err.Error()
		return response, err
	}
	response.Data = entity.Article{
		Title:       "Golang Project Structure",
		Description: "How to design your golang project structure",
		Timestamp: entity.Timestamp{
			CreateBy:   null.NewInt(0, true),
			CreateTime: null.NewTime(time.Now(), true),
		},
	}
	return response, nil
}

func (m *GetArticleModel) Validate(ctx context.Context) error {
	return nil
}
