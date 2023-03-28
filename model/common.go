package model

type CommonResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
