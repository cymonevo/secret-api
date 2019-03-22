package entity

import "time"

type BaseResponse struct {
	Status  int
	Message string
	Payload interface{}
}

type BaseTimestamp struct {
	CreateBy   int
	CreateTime time.Time
	UpdateBy   int
	UpdateTime time.Time
}
