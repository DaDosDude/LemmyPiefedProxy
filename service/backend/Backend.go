package backend

import (
	lemmyRequest "LemmyBeProxy/dto/request/lemmy"
	lemmyResponse "LemmyBeProxy/dto/response/lemmy"
	"LemmyBeProxy/http"
)

// Backend is the contract every backend implementation (Piefed, real Lemmy,
// and any future addition) must satisfy. Every method works entirely in
// Lemmy-shaped request/response types — that's already this proxy's
// canonical internal shape, since every controller was built to speak it
// regardless of what's actually behind it. A Piefed implementation
// translates to/from Piefed's real shape internally; a real-Lemmy
// implementation is close to a passthrough, since the shapes already
// match.
//
// This interface covers Post, Comment, Community, User, and Search
// methods so far, as the first five slices of a larger migration — Site
// and Upload still talk to the old *piefed.Piefed type directly and
// need the same treatment applied.
type Backend interface {
	GetPosts(request *lemmyRequest.GetPostsRequest, headers http.Headers) (*lemmyResponse.GetPostsResponse, error)
	GetPost(request *lemmyRequest.GetPostRequest, headers http.Headers) (*lemmyResponse.GetPostResponse, error)
	CreatePost(request *lemmyRequest.CreatePostRequest, headers http.Headers) (*lemmyResponse.GetPostResponse, error)
	EditPost(request *lemmyRequest.EditPostRequest, headers http.Headers) (*lemmyResponse.GetPostResponse, error)
	LikePost(request *lemmyRequest.CreatePostLikeRequest, headers http.Headers) (*lemmyResponse.GetPostResponse, error)
	MarkPostAsRead(request *lemmyRequest.MarkPostAsReadRequest, headers http.Headers) (*lemmyResponse.SuccessResponse, error)

	GetComments(request *lemmyRequest.GetCommentsRequest, headers http.Headers) (*lemmyResponse.GetCommentsResponse, error)
	GetComment(request *lemmyRequest.GetCommentRequest, headers http.Headers) (*lemmyResponse.GetCommentResponse, error)
	CreateComment(request *lemmyRequest.CreateCommentRequest, headers http.Headers) (*lemmyResponse.CreateCommentResponse, error)
	LikeComment(request *lemmyRequest.CreateCommentLikeRequest, headers http.Headers) (*lemmyResponse.GetCommentResponse, error)

	GetCommunity(request *lemmyRequest.GetCommunityRequest, headers http.Headers) (*lemmyResponse.GetCommunityResponse, error)
	GetCommunities(request *lemmyRequest.GetCommunitiesRequest, headers http.Headers) (*lemmyResponse.GetCommunitiesResponse, error)
	FollowCommunity(request *lemmyRequest.FollowCommunityRequest, headers http.Headers) (*lemmyResponse.CommunityResponse, error)
	BlockCommunity(request *lemmyRequest.BlockCommunityRequest, headers http.Headers) (*lemmyResponse.BlockCommunityResponse, error)

	Login(request *lemmyRequest.LoginRequest, headers http.Headers) (*lemmyResponse.LoginResponse, error)
	GetUnreadCount(headers http.Headers) (*lemmyResponse.GetUnreadCountResponse, error)
	GetUser(request *lemmyRequest.GetUserRequest, headers http.Headers) (*lemmyResponse.GetUserResponse, error)
	BlockPerson(request *lemmyRequest.BlockPersonRequest, headers http.Headers) (*lemmyResponse.BlockPersonResponse, error)
	SaveUserSettings(request *lemmyRequest.SaveUserSettingsRequest, headers http.Headers) (*lemmyResponse.SaveUserSettingsResponse, error)

	Search(request *lemmyRequest.SearchRequest, headers http.Headers) (*lemmyResponse.SearchResponse, error)
}
