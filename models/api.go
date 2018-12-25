package models

type LoginModel struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

type LoginResponseModel struct {
	SessionID string `json:"sessionId"`
}
