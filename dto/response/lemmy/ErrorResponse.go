package lemmy

import "LemmyBeProxy/dto/model/lemmy"

type ErrorResponse struct {
	Error   lemmy.ErrorCode `json:"error"`
	Message string          `json:"message,omitempty"`
}

func NewErrorResponse(error lemmy.ErrorCode) *ErrorResponse {
	return &ErrorResponse{
		Error: error,
	}
}

func NewErrorResponseWithMessage(error lemmy.ErrorCode, message string) *ErrorResponse {
	return &ErrorResponse{
		Error:   error,
		Message: message,
	}
}
