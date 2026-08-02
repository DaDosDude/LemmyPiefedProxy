package piefed

import (
	"LemmyPiefedApi/dto/request/piefed"
	piefedResponse "LemmyPiefedApi/dto/response/piefed"
	"LemmyPiefedApi/http"
	"LemmyPiefedApi/router"
)

func (receiver *Piefed) GetPosts(request *piefed.GetPostsRequest, headers http.Headers) (*piefedResponse.GetPostsResponse, error) {
	return defaultHandler[piefedResponse.GetPostsResponse](
		receiver,
		"/post/list",
		router.HttpMethodGet,
		request,
		headers,
	)
}

func (receiver *Piefed) GetPost(request *piefed.GetPostRequest, headers http.Headers) (*piefedResponse.GetPostResponse, error) {
	return defaultHandler[piefedResponse.GetPostResponse](
		receiver,
		"/post",
		router.HttpMethodGet,
		request,
		headers,
	)
}

func (receiver *Piefed) LikePost(request *piefed.LikePostRequest, headers http.Headers) (*piefedResponse.GetPostResponse, error) {
	return defaultHandler[piefedResponse.GetPostResponse](
		receiver,
		"/post/like",
		router.HttpMethodPost,
		request,
		headers,
	)
}

func (receiver *Piefed) MarkPostAsRead(request *piefed.MarkPostAsReadRequest, headers http.Headers) (*piefedResponse.SuccessResponse, error) {
	return defaultHandler[piefedResponse.SuccessResponse](
		receiver,
		"/post/mark_as_read",
		router.HttpMethodPost,
		request,
		headers,
	)
}
