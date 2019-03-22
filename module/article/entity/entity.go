package entity

import "github.com/cymon1997/go-backend/internal/entity"

type Article struct {
	Title       string
	Description string
	Content     string
	entity.BaseTimestamp
}
