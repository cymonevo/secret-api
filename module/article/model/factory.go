package model

import "github.com/cymon1997/go-backend/internal/entity"

type Factory interface {
	NewGetByIDModel() entity.BaseModel
}

type articleFactory struct {
}

func NewFactory() Factory {
	return &articleFactory{}
}

func (f *articleFactory) NewGetByIDModel() entity.BaseModel {
	return &getByIDModel{}
}
