package entity

type SecretData struct {
	ID    string `db:"id"`
	AppID string `db:"app_id"`
	Data  []byte `db:"data"`
	CreateTime
}

type InsertSecretRequest struct {
	AppID string `json:"app_id"`
	Data  []byte `json:"data"`
	//Data  json.RawMessage `json:"data"`
}

type InsertSecretResponse struct {
	Response
}
