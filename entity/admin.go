package entity

type AppData struct {
	AppID  string `db:"app_id"`
	Secret []byte `db:"secret,string"`
}

type RegisterRequest struct {
	AppID string `json:"app_id"`
}

type RegisterResponse struct {
	Secret []byte `json:"secret"`
	Response
}
