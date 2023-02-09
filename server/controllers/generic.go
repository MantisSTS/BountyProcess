package controllers

type HttpStdResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}
