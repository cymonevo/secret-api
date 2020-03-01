package entity

import "time"

type Timestamp struct {
	CreateTime time.Time `json:"create_time" db:"create_time"`
	CreateBy   string    `json:"create_by" db:"create_by"`
	UpdateTime time.Time `json:"update_time" db:"update_time"`
	UpdateBy   string    `json:"update_by" db:"update_by"`
}

type Response struct {
	Message string `json:"message,omitempty"`
	Status  int    `json:"status,omitempty"`
}
