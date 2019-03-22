package model

import (
	ientity "github.com/cymon1997/go-backend/internal/entity"
	"github.com/cymon1997/go-backend/module/article/entity"
	"time"
)

type getByIDModel struct {
}

func (g *getByIDModel) Call() (interface{}, error) {
	err := g.Validate()
	if err != nil {
		return nil, err
	}

	result := entity.Article{
		Title:       "Golang Project Structure",
		Description: "How to design your golang project structure",
		BaseTimestamp: ientity.BaseTimestamp{
			CreateBy:   1,
			CreateTime: time.Now(),
		},
	}
	return &result, nil
}

func (g *getByIDModel) Validate() error {
	return nil
}
