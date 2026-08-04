package frontend

import (
	lemmyRequest "LemmyBeProxy/dto/request/lemmy"
	lemmyResponse "LemmyBeProxy/dto/response/lemmy"
	"LemmyBeProxy/http"
)

// Frontend is the counterpart to backend.Backend, but for the other end
// of the pipeline: instead of choosing what this proxy talks to, it
// chooses what wire format this proxy accepts from and replies to
// clients with. Every method parses/builds using this proxy's canonical
// (current Lemmy 0.19.x-shaped) request/response types — the same ones
// backend.Backend works in — so a controller using both interfaces never
// needs to know or care which frontend or backend version is actually in
// play on either side of it.
//
// This interface only covers Post endpoints so far, mirroring how
// backend.Backend started with the same slice.
//
// Known gap: Frontend017's BuildSuccessResponse (used for mark_as_read)
// returns the canonical {success: bool} shape rather than the full
// PostResponse{post_view} real Lemmy 0.17.x actually returns there,
// since the canonical MarkPostAsRead call doesn't carry post_view data
// and building it accurately would need an extra backend round-trip —
// a real design decision deferred rather than rushed. Every other method
// in this slice is fully accurate to 0.17.2, confirmed against Lemmy's
// own source at that tag.
type Frontend interface {
	ParseGetPostsRequest(request *http.Request) (*lemmyRequest.GetPostsRequest, error)
	BuildGetPostsResponse(resp *lemmyResponse.GetPostsResponse) any

	ParseGetPostRequest(request *http.Request) (*lemmyRequest.GetPostRequest, error)
	BuildGetPostResponse(resp *lemmyResponse.GetPostResponse) any

	ParseCreatePostRequest(request *http.Request) (*lemmyRequest.CreatePostRequest, error)
	ParseEditPostRequest(request *http.Request) (*lemmyRequest.EditPostRequest, error)
	ParseCreatePostLikeRequest(request *http.Request) (*lemmyRequest.CreatePostLikeRequest, error)
	// BuildPostMutationResponse is shared by create/edit/like — the
	// canonical backend unifies all three into GetPostResponse, matching
	// how modern Lemmy actually behaves.
	BuildPostMutationResponse(resp *lemmyResponse.GetPostResponse) any

	ParseMarkPostAsReadRequest(request *http.Request) (*lemmyRequest.MarkPostAsReadRequest, error)
	BuildSuccessResponse(resp *lemmyResponse.SuccessResponse) any
}
