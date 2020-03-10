package entity

type AppData struct {
	AppID  string   `db:"app_id"`
	Secret [32]byte `db:"secret"`
}

type RegisterRequest struct {
	AppID string `json:"app_id"`
}

type RegisterResponse struct {
	Secret [32]byte `json:"secret"`
	Response
}
