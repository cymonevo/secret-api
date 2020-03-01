package model

import "context"

type BaseModel interface {
	Call(ctx context.Context) (interface{}, error)
	Validate(ctx context.Context) error
}
