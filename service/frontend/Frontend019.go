package frontend

import (
	lemmyRequest "LemmyBeProxy/dto/request/lemmy"
	lemmyResponse "LemmyBeProxy/dto/response/lemmy"
	"LemmyBeProxy/helper"
	"LemmyBeProxy/http"
)

// Frontend019 is the current-generation Lemmy wire format — this proxy's
// canonical shape already is this format, so every method here is a
// direct parse/return with no translation at all. This is exactly what
// every controller did inline before the Frontend interface existed.
type Frontend019 struct{}

func NewFrontend019() *Frontend019 {
	return &Frontend019{}
}

func (receiver *Frontend019) ParseGetPostsRequest(request *http.Request) (*lemmyRequest.GetPostsRequest, error) {
	return helper.ParseRequestQuery[lemmyRequest.GetPostsRequest](request)
}

func (receiver *Frontend019) BuildGetPostsResponse(resp *lemmyResponse.GetPostsResponse) any {
	return resp
}

func (receiver *Frontend019) ParseGetPostRequest(request *http.Request) (*lemmyRequest.GetPostRequest, error) {
	return helper.ParseRequestQuery[lemmyRequest.GetPostRequest](request)
}

func (receiver *Frontend019) BuildGetPostResponse(resp *lemmyResponse.GetPostResponse) any {
	return resp
}

func (receiver *Frontend019) ParseCreatePostRequest(request *http.Request) (*lemmyRequest.CreatePostRequest, error) {
	return helper.ParseRequest[lemmyRequest.CreatePostRequest](request)
}

func (receiver *Frontend019) ParseEditPostRequest(request *http.Request) (*lemmyRequest.EditPostRequest, error) {
	return helper.ParseRequest[lemmyRequest.EditPostRequest](request)
}

func (receiver *Frontend019) ParseCreatePostLikeRequest(request *http.Request) (*lemmyRequest.CreatePostLikeRequest, error) {
	return helper.ParseRequest[lemmyRequest.CreatePostLikeRequest](request)
}

func (receiver *Frontend019) BuildPostMutationResponse(resp *lemmyResponse.GetPostResponse) any {
	return resp
}

func (receiver *Frontend019) ParseMarkPostAsReadRequest(request *http.Request) (*lemmyRequest.MarkPostAsReadRequest, error) {
	return helper.ParseRequest[lemmyRequest.MarkPostAsReadRequest](request)
}

func (receiver *Frontend019) BuildSuccessResponse(resp *lemmyResponse.SuccessResponse) any {
	return resp
}
