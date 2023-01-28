package controllers

type ApiResponse struct {
	Status  uint   `json:"status"`
	Message string `json:"message"`
	Content any    `json:"content"`
	Ok      bool   `json:"ok"`
}

func NewSuccessResponse(content any) *ApiResponse {
	return &ApiResponse{
		Ok:      true,
		Message: "",
		Status:  200,
		Content: content,
	}
}

func NewErrorResponse(message string, statusCode uint) *ApiResponse {
	return &ApiResponse{
		Ok:      false,
		Message: message,
		Status:  statusCode,
		Content: nil,
	}
}
