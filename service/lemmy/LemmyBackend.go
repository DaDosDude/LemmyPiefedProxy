package lemmy

import (
	lemmyRequest "LemmyBeProxy/dto/request/lemmy"
	lemmyResponse "LemmyBeProxy/dto/response/lemmy"
	appHttp "LemmyBeProxy/http"
	"LemmyBeProxy/router"
)

// LemmyBackend implements the backend.Backend interface against a real
// Lemmy instance. Unlike PiefedBackend, this needs almost no translation:
// this proxy's canonical DTOs are already modeled on Lemmy's real field
// names, so requests and responses are forwarded close to as-is. The one
// exception is the Piefed-specific community_id + type_=Subscribed
// workaround in PiefedBackend, which doesn't apply here — that was a
// real Piefed backend bug, not a general Lemmy API quirk, so real Lemmy
// gets the request exactly as the client sent it.
type LemmyBackend struct {
	client *Lemmy
}

func NewLemmyBackend(client *Lemmy) *LemmyBackend {
	return &LemmyBackend{client: client}
}

func (receiver *LemmyBackend) GetPosts(request *lemmyRequest.GetPostsRequest, headers appHttp.Headers) (*lemmyResponse.GetPostsResponse, error) {
	return defaultHandler[lemmyResponse.GetPostsResponse](receiver.client, "/post/list", router.HttpMethodGet, request, headers)
}

func (receiver *LemmyBackend) GetPost(request *lemmyRequest.GetPostRequest, headers appHttp.Headers) (*lemmyResponse.GetPostResponse, error) {
	return defaultHandler[lemmyResponse.GetPostResponse](receiver.client, "/post", router.HttpMethodGet, request, headers)
}

func (receiver *LemmyBackend) CreatePost(request *lemmyRequest.CreatePostRequest, headers appHttp.Headers) (*lemmyResponse.GetPostResponse, error) {
	return defaultHandler[lemmyResponse.GetPostResponse](receiver.client, "/post", router.HttpMethodPost, request, headers)
}

func (receiver *LemmyBackend) EditPost(request *lemmyRequest.EditPostRequest, headers appHttp.Headers) (*lemmyResponse.GetPostResponse, error) {
	return defaultHandler[lemmyResponse.GetPostResponse](receiver.client, "/post", router.HttpMethodPut, request, headers)
}

func (receiver *LemmyBackend) LikePost(request *lemmyRequest.CreatePostLikeRequest, headers appHttp.Headers) (*lemmyResponse.GetPostResponse, error) {
	return defaultHandler[lemmyResponse.GetPostResponse](receiver.client, "/post/like", router.HttpMethodPost, request, headers)
}

func (receiver *LemmyBackend) MarkPostAsRead(request *lemmyRequest.MarkPostAsReadRequest, headers appHttp.Headers) (*lemmyResponse.SuccessResponse, error) {
	return defaultHandler[lemmyResponse.SuccessResponse](receiver.client, "/post/mark_as_read", router.HttpMethodPost, request, headers)
}
