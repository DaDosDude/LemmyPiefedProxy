package frontend

import (
	lemmyRequest "LemmyBeProxy/dto/request/lemmy"
	lemmyRequest017 "LemmyBeProxy/dto/request/lemmy017"
	lemmyResponse "LemmyBeProxy/dto/response/lemmy"
	lemmyResponse017 "LemmyBeProxy/dto/response/lemmy017"
	"LemmyBeProxy/helper"
	"LemmyBeProxy/http"
)

// Frontend017 implements the Frontend interface for Lemmy 0.17.x clients
// (built for lemmyBB, which is pinned to lemmy_api_common 0.17.2). Every
// shape difference here is confirmed against Lemmy's own source at tag
// 0.17.2, not assumed — see dto/request/lemmy017 and
// dto/response/lemmy017 for the specific differences each type documents.
type Frontend017 struct{}

func NewFrontend017() *Frontend017 {
	return &Frontend017{}
}

func (receiver *Frontend017) ParseGetPostsRequest(request *http.Request) (*lemmyRequest.GetPostsRequest, error) {
	reqDto, err := helper.ParseRequestQuery[lemmyRequest017.GetPostsRequest](request)
	if err != nil {
		return nil, err
	}

	return &lemmyRequest.GetPostsRequest{
		Type:          reqDto.Type,
		Sort:          reqDto.Sort,
		Page:          reqDto.Page,
		Limit:         reqDto.Limit,
		CommunityId:   reqDto.CommunityId,
		CommunityName: reqDto.CommunityName,
		SavedOnly:     reqDto.SavedOnly,
		// LikedOnly, DislikedOnly, PageCursor, ShowHidden, ShowNsfw,
		// ShowRead don't exist in 0.17.x — left unset.
	}, nil
}

func (receiver *Frontend017) BuildGetPostsResponse(resp *lemmyResponse.GetPostsResponse) any {
	// 0.17.x has no next_page concept at all — dropped, not translated.
	return &lemmyResponse017.GetPostsResponse{
		Posts: resp.Posts,
	}
}

func (receiver *Frontend017) ParseGetPostRequest(request *http.Request) (*lemmyRequest.GetPostRequest, error) {
	reqDto, err := helper.ParseRequestQuery[lemmyRequest017.GetPostRequest](request)
	if err != nil {
		return nil, err
	}

	return &lemmyRequest.GetPostRequest{
		Id:        reqDto.Id,
		CommentId: reqDto.CommentId,
	}, nil
}

func (receiver *Frontend017) BuildGetPostResponse(resp *lemmyResponse.GetPostResponse) any {
	// 0.17.x has no cross_posts field and expects an "online" viewer
	// count our canonical model doesn't track at all — 0 is the honest
	// value here, not a guess standing in for real data.
	return &lemmyResponse017.GetPostResponse{
		PostView:      resp.PostView,
		CommunityView: resp.CommunityView,
		Moderators:    resp.Moderators,
		Online:        0,
	}
}

func (receiver *Frontend017) ParseCreatePostRequest(request *http.Request) (*lemmyRequest.CreatePostRequest, error) {
	reqDto, err := helper.ParseRequest[lemmyRequest017.CreatePostRequest](request)
	if err != nil {
		return nil, err
	}

	return &lemmyRequest.CreatePostRequest{
		Name:        reqDto.Name,
		CommunityId: reqDto.CommunityId,
		Body:        reqDto.Body,
		Url:         reqDto.Url,
		Nsfw:        reqDto.Nsfw,
		LanguageId:  reqDto.LanguageId,
		Honeypot:    reqDto.Honeypot,
	}, nil
}

func (receiver *Frontend017) ParseEditPostRequest(request *http.Request) (*lemmyRequest.EditPostRequest, error) {
	reqDto, err := helper.ParseRequest[lemmyRequest017.EditPostRequest](request)
	if err != nil {
		return nil, err
	}

	return &lemmyRequest.EditPostRequest{
		PostId:     reqDto.PostId,
		Name:       reqDto.Name,
		Body:       reqDto.Body,
		Url:        reqDto.Url,
		Nsfw:       reqDto.Nsfw,
		LanguageId: reqDto.LanguageId,
	}, nil
}

func (receiver *Frontend017) ParseCreatePostLikeRequest(request *http.Request) (*lemmyRequest.CreatePostLikeRequest, error) {
	reqDto, err := helper.ParseRequest[lemmyRequest017.CreatePostLikeRequest](request)
	if err != nil {
		return nil, err
	}

	return &lemmyRequest.CreatePostLikeRequest{
		PostId: reqDto.PostId,
		Score:  reqDto.Score,
	}, nil
}

func (receiver *Frontend017) BuildPostMutationResponse(resp *lemmyResponse.GetPostResponse) any {
	// 0.17.x's create/edit/like all return the lean PostResponse{post_view}
	// shape — confirmed against each handler in Lemmy's own source, not
	// the fuller GetPostResponse used for fetching.
	return &lemmyResponse017.PostResponse{
		PostView: resp.PostView,
	}
}

func (receiver *Frontend017) ParseMarkPostAsReadRequest(request *http.Request) (*lemmyRequest.MarkPostAsReadRequest, error) {
	reqDto, err := helper.ParseRequest[lemmyRequest017.MarkPostAsReadRequest](request)
	if err != nil {
		return nil, err
	}

	return &lemmyRequest.MarkPostAsReadRequest{
		PostId: helper.ToPointer(reqDto.PostId),
		Read:   reqDto.Read,
	}, nil
}

// BuildSuccessResponse is the one known gap in this slice — see the
// comment on the Frontend interface for why. Real Lemmy 0.17.x expects
// PostResponse{post_view} here; this returns the canonical
// {success: bool} shape instead, since building an accurate post_view
// would need an extra backend round-trip that's a real design decision,
// not something to bolt on silently.
func (receiver *Frontend017) BuildSuccessResponse(resp *lemmyResponse.SuccessResponse) any {
	return resp
}
