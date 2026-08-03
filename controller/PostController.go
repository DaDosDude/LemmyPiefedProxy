package controller

import (
	lemmyModel "LemmyPiefedApi/dto/model/lemmy"
	piefedModel "LemmyPiefedApi/dto/model/piefed"
	"LemmyPiefedApi/dto/request/lemmy"
	"LemmyPiefedApi/dto/request/piefed"
	lemmyResponse "LemmyPiefedApi/dto/response/lemmy"
	"LemmyPiefedApi/helper"
	"LemmyPiefedApi/helper/converter"
	"LemmyPiefedApi/http"
	pfService "LemmyPiefedApi/service/piefed"
	goHttp "net/http"
)

type PostController struct {
	piefed *pfService.Piefed
}

func NewPostController(piefed *pfService.Piefed) *PostController {
	return &PostController{
		piefed: piefed,
	}
}

func (receiver *PostController) GetPosts(request *http.Request) (*http.Response, error) {
	reqDto, err := helper.ParseRequestQuery[lemmy.GetPostsRequest](request)
	if err != nil {
		return helper.ConvertValidationErrorsToResponse(err), nil
	}

	if reqDto.PageCursor != nil {
		return http.NotImplementedFeature(
			"Cursor based navigation is not available",
		), nil
	}

	// Piefed's /post/list ignores community_id/community_name entirely when
	// type_ is also set to Subscribed (confirmed directly: sending both
	// returns the full subscribed feed, not the requested community's
	// posts — a real Piefed backend quirk, not a translation gap here).
	// A community-scoped request is unambiguous on its own; the listing
	// type serves no purpose once a specific community is named, so it's
	// dropped whenever either community field is present rather than
	// risking Piefed silently discarding the more specific filter.
	isCommunityScoped := reqDto.CommunityId != nil || reqDto.CommunityName != nil

	resp, err := receiver.piefed.GetPosts(&piefed.GetPostsRequest{
		Type: helper.SafeDereference(reqDto.Type, func(in lemmyModel.ListingType) *piefedModel.ListingType {
			if isCommunityScoped {
				return nil
			}
			return helper.ToPointer(converter.ReverseConvertListingType(in))
		}),
		Sort: helper.SafeDereference(reqDto.Sort, func(in lemmyModel.SortType) *piefedModel.SortType {
			return helper.ToPointer(converter.ReverseConvertSortType(in))
		}),
		Page:          reqDto.Page,
		Limit:         reqDto.Limit,
		CommunityId:   reqDto.CommunityId,
		PersonId:      nil,
		CommunityName: reqDto.CommunityName,
		LikedOnly:     reqDto.LikedOnly,
	}, request.Headers)
	if err != nil {
		return nil, err
	}

	return &http.Response{
		StatusCode: goHttp.StatusOK,
		Body: &lemmyResponse.GetPostsResponse{
			NextPage: resp.NextPage,
			Posts:    helper.MapSlice(resp.Posts, converter.ConvertPostView),
		},
	}, nil
}

func (receiver *PostController) GetPost(request *http.Request) (*http.Response, error) {
	reqDto, err := helper.ParseRequestQuery[lemmy.GetPostRequest](request)
	if err != nil {
		return helper.ConvertValidationErrorsToResponse(err), nil
	}

	resp, err := receiver.piefed.GetPost(&piefed.GetPostRequest{
		CommentId: reqDto.CommentId,
		Id:        reqDto.Id,
	}, request.Headers)
	if err != nil {
		return nil, err
	}

	return &http.Response{
		StatusCode: goHttp.StatusOK,
		Body: &lemmyResponse.GetPostResponse{
			CommunityView: converter.ConvertCommunityView(resp.CommunityView),
			CrossPosts:    helper.MapSlice(resp.CrossPosts, converter.ConvertPostView),
			Moderators:    helper.MapSlice(resp.Moderators, converter.ConvertCommunityModeratorView),
			PostView:      converter.ConvertPostView(resp.PostView),
		},
	}, nil
}

func (receiver *PostController) MarkPostAsRead(request *http.Request) (*http.Response, error) {
	reqDto, err := helper.ParseRequest[lemmy.MarkPostAsReadRequest](request)
	if err != nil {
		return helper.ConvertValidationErrorsToResponse(err), nil
	}

	resp, err := receiver.piefed.MarkPostAsRead(&piefed.MarkPostAsReadRequest{
		PostId:  reqDto.PostId,
		PostIds: reqDto.PostIds,
		Read:    reqDto.Read,
	}, request.Headers)
	if err != nil {
		return nil, err
	}

	return &http.Response{
		StatusCode: goHttp.StatusOK,
		Body: &lemmyResponse.SuccessResponse{
			Success: resp.Success,
		},
	}, nil
}

// CreatePost maps Lemmy's Name field to Piefed's Title — same content,
// different field name. Honeypot has no Piefed equivalent and is dropped.
func (receiver *PostController) CreatePost(request *http.Request) (*http.Response, error) {
	reqDto, err := helper.ParseRequest[lemmy.CreatePostRequest](request)
	if err != nil {
		return helper.ConvertValidationErrorsToResponse(err), nil
	}

	resp, err := receiver.piefed.CreatePost(&piefed.CreatePostRequest{
		Title:       reqDto.Name,
		CommunityId: reqDto.CommunityId,
		Body:        reqDto.Body,
		Url:         reqDto.Url,
		Nsfw:        reqDto.Nsfw,
		LanguageId:  reqDto.LanguageId,
	}, request.Headers)
	if err != nil {
		return nil, err
	}

	return &http.Response{
		StatusCode: goHttp.StatusOK,
		Body: &lemmyResponse.GetPostResponse{
			CommunityView: converter.ConvertCommunityView(resp.CommunityView),
			CrossPosts:    helper.MapSlice(resp.CrossPosts, converter.ConvertPostView),
			Moderators:    helper.MapSlice(resp.Moderators, converter.ConvertCommunityModeratorView),
			PostView:      converter.ConvertPostView(resp.PostView),
		},
	}, nil
}

func (receiver *PostController) EditPost(request *http.Request) (*http.Response, error) {
	reqDto, err := helper.ParseRequest[lemmy.EditPostRequest](request)
	if err != nil {
		return helper.ConvertValidationErrorsToResponse(err), nil
	}

	resp, err := receiver.piefed.EditPost(&piefed.EditPostRequest{
		PostId:     reqDto.PostId,
		Title:      reqDto.Name,
		Body:       reqDto.Body,
		Url:        reqDto.Url,
		Nsfw:       reqDto.Nsfw,
		LanguageId: reqDto.LanguageId,
	}, request.Headers)
	if err != nil {
		return nil, err
	}

	return &http.Response{
		StatusCode: goHttp.StatusOK,
		Body: &lemmyResponse.GetPostResponse{
			CommunityView: converter.ConvertCommunityView(resp.CommunityView),
			CrossPosts:    helper.MapSlice(resp.CrossPosts, converter.ConvertPostView),
			Moderators:    helper.MapSlice(resp.Moderators, converter.ConvertCommunityModeratorView),
			PostView:      converter.ConvertPostView(resp.PostView),
		},
	}, nil
}

func (receiver *PostController) LikePost(request *http.Request) (*http.Response, error) {
	reqDto, err := helper.ParseRequest[lemmy.CreatePostLikeRequest](request)
	if err != nil {
		return helper.ConvertValidationErrorsToResponse(err), nil
	}

	resp, err := receiver.piefed.LikePost(&piefed.LikePostRequest{
		PostId: reqDto.PostId,
		Score:  reqDto.Score,
	}, request.Headers)
	if err != nil {
		return nil, err
	}

	return &http.Response{
		StatusCode: goHttp.StatusOK,
		Body: &lemmyResponse.GetPostResponse{
			CommunityView: converter.ConvertCommunityView(resp.CommunityView),
			CrossPosts:    helper.MapSlice(resp.CrossPosts, converter.ConvertPostView),
			Moderators:    helper.MapSlice(resp.Moderators, converter.ConvertCommunityModeratorView),
			PostView:      converter.ConvertPostView(resp.PostView),
		},
	}, nil
}
