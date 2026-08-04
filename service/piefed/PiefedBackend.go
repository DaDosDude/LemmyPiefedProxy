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
