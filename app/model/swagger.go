package model

type ProductGetAllSuccessResponse struct {
	Message string            `json:"message"`
	Data    []ProductResponse `json:"data"`
}

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}
