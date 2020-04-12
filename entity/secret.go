package entity

type SecretData struct {
	ID    string `db:"id"`
	AppID string `db:"app_id"`
	Data  []byte `db:"data"`
	CreateTime
}

type InsertSecretRequest struct {
	AppID string      `json:"app_id"`
	Data  interface{} `json:"data"`
}

type InsertSecretResponse struct {
	Response
}

type GetSecretRequest struct {
	AppID string `json:"app_id"`
}

type GetSecretResponse struct {
	Data interface{} `json:"data,omitempty"`
	Response
}
