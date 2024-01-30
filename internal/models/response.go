package models

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
	Error   string `json:"error"`
}
