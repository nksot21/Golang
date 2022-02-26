package models

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Error   bool        `json:"error"`
	Data    interface{} `json:"data"`
}
