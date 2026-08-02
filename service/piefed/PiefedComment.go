package piefed

import (
	"LemmyPiefedApi/dto/request/piefed"
	piefedResponse "LemmyPiefedApi/dto/response/piefed"
	"LemmyPiefedApi/http"
	"LemmyPiefedApi/router"
)

func (receiver *Piefed) GetComments(request *piefed.GetCommentsRequest, headers http.Headers) (*piefedResponse.GetCommentsResponse, error) {
	return defaultHandler[piefedResponse.GetCommentsResponse](
		receiver,
		"/comment/list",
		router.HttpMethodGet,
		request,
		headers,
	)
}

func (receiver *Piefed) GetComment(request *piefed.GetCommentRequest, headers http.Headers) (*piefedResponse.GetCommentResponse, error) {
	return defaultHandler[piefedResponse.GetCommentResponse](
		receiver,
		"/comment",
		router.HttpMethodGet,
		request,
		headers,
	)
}

func (receiver *Piefed) CreateComment(request *piefed.CreateCommentRequest, headers http.Headers) (*piefedResponse.CreateCommentResponse, error) {
	return defaultHandler[piefedResponse.CreateCommentResponse](
		receiver,
		"/comment",
		router.HttpMethodPost,
		request,
		headers,
	)
}

func (receiver *Piefed) LikeComment(request *piefed.LikeCommentRequest, headers http.Headers) (*piefedResponse.GetCommentResponse, error) {
	return defaultHandler[piefedResponse.GetCommentResponse](
		receiver,
		"/comment/like",
		router.HttpMethodPost,
		request,
		headers,
	)
}
