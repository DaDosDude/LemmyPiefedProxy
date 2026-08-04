package helper

import (
	lemmyModel "LemmyBeProxy/dto/model/lemmy"
	"LemmyBeProxy/dto/response/lemmy"
	"LemmyBeProxy/dto/response/piefed"
)

func ConvertPiefedErrorToLemmyError(err *piefed.ErrorResponse) *lemmy.ErrorResponse {
	var lemmyErrorCode lemmyModel.ErrorCode
	var message string

	switch err.ErrorCode {
	case "incorrect_login":
		lemmyErrorCode = lemmyModel.ErrorCodeIncorrectLogin
		break
	default:
		return &lemmy.ErrorResponse{
			Error:   lemmyModel.ErrorCodeUnknown,
			Message: "Piefed error: " + err.ErrorCode,
		}
	}

	return &lemmy.ErrorResponse{
		Error:   lemmyErrorCode,
		Message: message,
	}
}
