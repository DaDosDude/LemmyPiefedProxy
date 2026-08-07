package piefed

import (
	lemmyModel "LemmyBeProxy/dto/model/lemmy"
	piefedModel "LemmyBeProxy/dto/model/piefed"
	lemmyRequest "LemmyBeProxy/dto/request/lemmy"
	piefedRequest "LemmyBeProxy/dto/request/piefed"
	lemmyResponse "LemmyBeProxy/dto/response/lemmy"
	piefedResponse "LemmyBeProxy/dto/response/piefed"
	"LemmyBeProxy/helper"
	"LemmyBeProxy/helper/converter"
	appHttp "LemmyBeProxy/http"
)

// PiefedBackend implements the backend.Backend interface, translating
// between Lemmy-shaped requests/responses (this proxy's canonical shape)
// and Piefed's real API shape. It wraps the low-level *Piefed client,
// which just makes the raw HTTP calls.
type PiefedBackend struct {
	client *Piefed
}

func NewPiefedBackend(client *Piefed) *PiefedBackend {
	return &PiefedBackend{client: client}
}

// convertGetPostResponse is shared by every Post method that returns
// Piefed's GetPostResponse shape (get, create, edit, like all do).
func convertGetPostResponse(resp *piefedResponse.GetPostResponse) *lemmyResponse.GetPostResponse {
	return &lemmyResponse.GetPostResponse{
		CommunityView: converter.ConvertCommunityView(resp.CommunityView),
		CrossPosts:    helper.MapSlice(resp.CrossPosts, converter.ConvertPostView),
		Moderators:    helper.MapSlice(resp.Moderators, converter.ConvertCommunityModeratorView),
		PostView:      converter.ConvertPostView(resp.PostView),
	}
}

func (receiver *PiefedBackend) GetPosts(request *lemmyRequest.GetPostsRequest, headers appHttp.Headers) (*lemmyResponse.GetPostsResponse, error) {
	// Piefed's /post/list ignores community_id/community_name entirely
	// when type_ is also set to Subscribed (confirmed directly: sending
	// both returns the full subscribed feed, not the requested
	// community's posts — a real Piefed backend quirk). A
	// community-scoped request is unambiguous on its own, so the listing
	// type is dropped whenever either community field is present rather
	// than risking Piefed silently discarding the more specific filter.
	isCommunityScoped := request.CommunityId != nil || request.CommunityName != nil

	resp, err := receiver.client.GetPosts(&piefedRequest.GetPostsRequest{
		Type: helper.SafeDereference(request.Type, func(in lemmyModel.ListingType) *piefedModel.ListingType {
			if isCommunityScoped {
				return nil
			}
			return helper.ToPointer(converter.ReverseConvertListingType(in))
		}),
		Sort: helper.SafeDereference(request.Sort, func(in lemmyModel.SortType) *piefedModel.SortType {
			return helper.ToPointer(converter.ReverseConvertSortType(in))
		}),
		Page:          request.Page,
		Limit:         request.Limit,
		CommunityId:   request.CommunityId,
		PersonId:      nil,
		CommunityName: request.CommunityName,
		LikedOnly:     request.LikedOnly,
	}, headers)
	if err != nil {
		return nil, err
	}

	return &lemmyResponse.GetPostsResponse{
		NextPage: resp.NextPage,
		Posts:    helper.MapSlice(resp.Posts, converter.ConvertPostView),
	}, nil
}

func (receiver *PiefedBackend) GetPost(request *lemmyRequest.GetPostRequest, headers appHttp.Headers) (*lemmyResponse.GetPostResponse, error) {
	resp, err := receiver.client.GetPost(&piefedRequest.GetPostRequest{
		CommentId: request.CommentId,
		Id:        request.Id,
	}, headers)
	if err != nil {
		return nil, err
	}

	return convertGetPostResponse(resp), nil
}

// CreatePost maps Lemmy's Name field to Piefed's Title — same content,
// different field name. Honeypot has no Piefed equivalent and is dropped.
func (receiver *PiefedBackend) CreatePost(request *lemmyRequest.CreatePostRequest, headers appHttp.Headers) (*lemmyResponse.GetPostResponse, error) {
	resp, err := receiver.client.CreatePost(&piefedRequest.CreatePostRequest{
		Title:       request.Name,
		CommunityId: request.CommunityId,
		Body:        request.Body,
		Url:         request.Url,
		Nsfw:        request.Nsfw,
		LanguageId:  request.LanguageId,
	}, headers)
	if err != nil {
		return nil, err
	}

	return convertGetPostResponse(resp), nil
}

func (receiver *PiefedBackend) EditPost(request *lemmyRequest.EditPostRequest, headers appHttp.Headers) (*lemmyResponse.GetPostResponse, error) {
	resp, err := receiver.client.EditPost(&piefedRequest.EditPostRequest{
		PostId:     request.PostId,
		Title:      request.Name,
		Body:       request.Body,
		Url:        request.Url,
		Nsfw:       request.Nsfw,
		LanguageId: request.LanguageId,
	}, headers)
	if err != nil {
		return nil, err
	}

	return convertGetPostResponse(resp), nil
}

func (receiver *PiefedBackend) LikePost(request *lemmyRequest.CreatePostLikeRequest, headers appHttp.Headers) (*lemmyResponse.GetPostResponse, error) {
	resp, err := receiver.client.LikePost(&piefedRequest.LikePostRequest{
		PostId: request.PostId,
		Score:  request.Score,
	}, headers)
	if err != nil {
		return nil, err
	}

	return convertGetPostResponse(resp), nil
}

func (receiver *PiefedBackend) MarkPostAsRead(request *lemmyRequest.MarkPostAsReadRequest, headers appHttp.Headers) (*lemmyResponse.SuccessResponse, error) {
	resp, err := receiver.client.MarkPostAsRead(&piefedRequest.MarkPostAsReadRequest{
		PostId:  request.PostId,
		PostIds: request.PostIds,
		Read:    request.Read,
	}, headers)
	if err != nil {
		return nil, err
	}

	return &lemmyResponse.SuccessResponse{Success: resp.Success}, nil
}

func (receiver *PiefedBackend) GetComments(request *lemmyRequest.GetCommentsRequest, headers appHttp.Headers) (*lemmyResponse.GetCommentsResponse, error) {
	resp, err := receiver.client.GetComments(&piefedRequest.GetCommentsRequest{
		Type: helper.SafeDereference(request.Type, func(in lemmyModel.ListingType) *piefedModel.ListingType {
			return helper.ToPointer(converter.ReverseConvertListingType(in))
		}),
		PersonId:    nil,
		MaxDepth:    request.MaxDepth,
		Page:        request.Page,
		ParentId:    request.ParentId,
		CommunityId: request.CommunityId,
		PostId:      request.PostId,
		Limit:       request.Limit,
		Sort: helper.SafeDereference(request.Sort, func(in lemmyModel.CommentSortType) *piefedModel.CommentSortType {
			return helper.ToPointer(converter.ReverseConvertCommentSortType(in))
		}),
	}, headers)
	if err != nil {
		return nil, err
	}

	return &lemmyResponse.GetCommentsResponse{
		Comments: helper.MapSlice(resp.Comments, converter.ConvertCommentView),
	}, nil
}

func (receiver *PiefedBackend) GetComment(request *lemmyRequest.GetCommentRequest, headers appHttp.Headers) (*lemmyResponse.GetCommentResponse, error) {
	resp, err := receiver.client.GetComment(&piefedRequest.GetCommentRequest{
		Id: request.Id,
	}, headers)
	if err != nil {
		return nil, err
	}

	return &lemmyResponse.GetCommentResponse{
		CommentView:  converter.ConvertCommentView(resp.CommentView),
		RecipientIds: []uint{},
	}, nil
}

func (receiver *PiefedBackend) CreateComment(request *lemmyRequest.CreateCommentRequest, headers appHttp.Headers) (*lemmyResponse.CreateCommentResponse, error) {
	resp, err := receiver.client.CreateComment(&piefedRequest.CreateCommentRequest{
		Body:       request.Content,
		PostId:     request.PostId,
		ParentId:   request.ParentId,
		LanguageId: request.LanguageId,
	}, headers)
	if err != nil {
		return nil, err
	}

	return &lemmyResponse.CreateCommentResponse{
		CommentView:  converter.ConvertCommentView(resp.CommentView),
		RecipientIds: []uint{},
	}, nil
}

func (receiver *PiefedBackend) LikeComment(request *lemmyRequest.CreateCommentLikeRequest, headers appHttp.Headers) (*lemmyResponse.GetCommentResponse, error) {
	resp, err := receiver.client.LikeComment(&piefedRequest.LikeCommentRequest{
		CommentId: request.CommentId,
		Score:     request.Score,
	}, headers)
	if err != nil {
		return nil, err
	}

	return &lemmyResponse.GetCommentResponse{
		CommentView:  converter.ConvertCommentView(resp.CommentView),
		RecipientIds: []uint{},
	}, nil
}

func (receiver *PiefedBackend) GetCommunity(request *lemmyRequest.GetCommunityRequest, headers appHttp.Headers) (*lemmyResponse.GetCommunityResponse, error) {
	resp, err := receiver.client.GetCommunity(&piefedRequest.GetCommunityRequest{
		Id:   request.Id,
		Name: request.Name,
	}, headers)
	if err != nil {
		return nil, err
	}

	return &lemmyResponse.GetCommunityResponse{
		CommunityView: converter.ConvertCommunityView(resp.CommunityView),
		Moderators:    helper.MapSlice(resp.Moderators, converter.ConvertCommunityModeratorView),
	}, nil
}

func (receiver *PiefedBackend) GetCommunities(request *lemmyRequest.GetCommunitiesRequest, headers appHttp.Headers) (*lemmyResponse.GetCommunitiesResponse, error) {
	resp, err := receiver.client.GetCommunities(&piefedRequest.GetCommunitiesRequest{
		Type: helper.SafeDereference(request.Type, func(in lemmyModel.ListingType) *piefedModel.ListingType {
			return helper.ToPointer(converter.ReverseConvertListingType(in))
		}),
		Sort: helper.SafeDereference(request.Sort, func(in lemmyModel.SortType) *piefedModel.SortType {
			return helper.ToPointer(converter.ReverseConvertCommunitySortType(in))
		}),
		ShowNsfw: request.ShowNsfw,
		Page:     request.Page,
		Limit:    request.Limit,
	}, headers)
	if err != nil {
		return nil, err
	}

	return &lemmyResponse.GetCommunitiesResponse{
		Communities: helper.MapSlice(resp.Communities, converter.ConvertCommunityView),
	}, nil
}

func (receiver *PiefedBackend) FollowCommunity(request *lemmyRequest.FollowCommunityRequest, headers appHttp.Headers) (*lemmyResponse.CommunityResponse, error) {
	resp, err := receiver.client.FollowCommunity(&piefedRequest.FollowCommunityRequest{
		CommunityId: request.CommunityId,
		Follow:      request.Follow,
	}, headers)
	if err != nil {
		return nil, err
	}

	return &lemmyResponse.CommunityResponse{
		CommunityView:       converter.ConvertCommunityView(resp.CommunityView),
		DiscussionLanguages: resp.DiscussionLanguages,
	}, nil
}

func (receiver *PiefedBackend) BlockCommunity(request *lemmyRequest.BlockCommunityRequest, headers appHttp.Headers) (*lemmyResponse.BlockCommunityResponse, error) {
	resp, err := receiver.client.BlockCommunity(&piefedRequest.BlockCommunityRequest{
		CommunityId: request.CommunityId,
		Block:       request.Block,
	}, headers)
	if err != nil {
		return nil, err
	}

	return &lemmyResponse.BlockCommunityResponse{
		Blocked:       resp.Blocked,
		CommunityView: converter.ConvertCommunityView(resp.CommunityView),
	}, nil
}
