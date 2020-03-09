package entity

import (
	"gopkg.in/guregu/null.v3"
)

type CreateTime struct {
	CreateTime null.Time `json:"create_time" db:"create_time"`
	CreateBy   null.Int  `json:"create_by" db:"create_by"`
}

type Timestamp struct {
	CreateTime null.Time `json:"create_time" db:"create_time"`
	CreateBy   null.Int  `json:"create_by" db:"create_by"`
	UpdateTime null.Time `json:"update_time" db:"update_time"`
	UpdateBy   null.Int  `json:"update_by" db:"update_by"`
}

type Response struct {
	Message string `json:"message,omitempty"`
	Status  int    `json:"status,omitempty"`
}
