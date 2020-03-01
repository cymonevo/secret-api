package model

import (
	"context"
	"time"

	base "github.com/cymon1997/go-backend/internal/base/entity"
	"github.com/cymon1997/go-backend/internal/base/model"
	"github.com/cymon1997/go-backend/module/article/entity"
)

type getByIDModel struct {
	model.BaseModel
}

func (m *getByIDModel) Call(ctx context.Context) (interface{}, error) {
	err := m.Validate(ctx)
	if err != nil {
		return nil, err
	}
	result := entity.Article{
		Title:       "Golang Project Structure",
		Description: "How to design your golang project structure",
		Timestamp: base.Timestamp{
			CreateBy:   "",
			CreateTime: time.Now(),
		},
	}
	return &result, nil
}

func (m *getByIDModel) Validate(ctx context.Context) error {
	return nil
}
