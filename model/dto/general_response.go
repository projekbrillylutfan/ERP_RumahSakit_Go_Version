package dto

type GeneralResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type GeneralResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   any    `json:"error"`
}