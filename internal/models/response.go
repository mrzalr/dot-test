package models

import "net/http"

type response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
	Error   string `json:"error"`
}

func NewResponseCreated(data any) *response {
	return &response{
		Status:  http.StatusCreated,
		Message: "created",
		Data:    data,
		Error:   "",
	}
}
