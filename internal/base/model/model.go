package model

import "context"

type Model interface {
	Call(ctx context.Context) (interface{}, error)
	Validate(ctx context.Context) error
}
