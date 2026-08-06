package controller

import (
	lemmyModel "LemmyBeProxy/dto/model/lemmy"
	piefedModel "LemmyBeProxy/dto/model/piefed"
	"LemmyBeProxy/dto/request/piefed"
	lemmyResponse "LemmyBeProxy/dto/response/lemmy"
	"LemmyBeProxy/helper"
	"LemmyBeProxy/helper/converter"
	"LemmyBeProxy/http"
	"LemmyBeProxy/service/frontend"
	pfService "LemmyBeProxy/service/piefed"
	goHttp "net/http"
)

// CommunityController now uses frontend.Frontend for parsing/building
// (wire-format-agnostic), but still calls the raw Piefed client directly
// for the actual backend call — Community hasn't been migrated onto
// backend.Backend yet, only Post and Comment have. The two interfaces
// are independent axes; this controller is on the Frontend axis only
// for now.
type CommunityController struct {
	piefed   *pfService.Piefed
	frontend frontend.Frontend
}

func NewCommunityController(piefed *pfService.Piefed, frontend frontend.Frontend) *CommunityController {
	return &CommunityController{
		piefed:   piefed,
		frontend: frontend,
	}
}

func (receiver *CommunityController) GetCommunity(request *http.Request) (*http.Response, error) {
	reqDto, err := receiver.frontend.ParseGetCommunityRequest(request)
	if err != nil {
		return helper.ConvertValidationErrorsToResponse(err), nil
	}

	resp, err := receiver.piefed.GetCommunity(&piefed.GetCommunityRequest{
		Id:   reqDto.Id,
		Name: reqDto.Name,
	}, request.Headers)
	if err != nil {
		return nil, err
	}

	canonical := &lemmyResponse.GetCommunityResponse{
		CommunityView: converter.ConvertCommunityView(resp.CommunityView),
		Moderators:    helper.MapSlice(resp.Moderators, converter.ConvertCommunityModeratorView),
	}

	return &http.Response{StatusCode: goHttp.StatusOK, Body: receiver.frontend.BuildGetCommunityResponse(canonical)}, nil
}

func (receiver *CommunityController) GetCommunities(request *http.Request) (*http.Response, error) {
	reqDto, err := receiver.frontend.ParseGetCommunitiesRequest(request)
	if err != nil {
		return helper.ConvertValidationErrorsToResponse(err), nil
	}

	resp, err := receiver.piefed.GetCommunities(&piefed.GetCommunitiesRequest{
		Type: helper.SafeDereference(reqDto.Type, func(in lemmyModel.ListingType) *piefedModel.ListingType {
			return helper.ToPointer(converter.ReverseConvertListingType(in))
		}),
		Sort: helper.SafeDereference(reqDto.Sort, func(in lemmyModel.SortType) *piefedModel.SortType {
			return helper.ToPointer(converter.ReverseConvertCommunitySortType(in))
		}),
		ShowNsfw: reqDto.ShowNsfw,
		Page:     reqDto.Page,
		Limit:    reqDto.Limit,
	}, request.Headers)
	if err != nil {
		return nil, err
	}

	canonical := &lemmyResponse.GetCommunitiesResponse{
		Communities: helper.MapSlice(resp.Communities, converter.ConvertCommunityView),
	}

	return &http.Response{StatusCode: goHttp.StatusOK, Body: receiver.frontend.BuildGetCommunitiesResponse(canonical)}, nil
}

func (receiver *CommunityController) FollowCommunity(request *http.Request) (*http.Response, error) {
	reqDto, err := receiver.frontend.ParseFollowCommunityRequest(request)
	if err != nil {
		return helper.ConvertValidationErrorsToResponse(err), nil
	}

	resp, err := receiver.piefed.FollowCommunity(&piefed.FollowCommunityRequest{
		CommunityId: reqDto.CommunityId,
		Follow:      reqDto.Follow,
	}, request.Headers)
	if err != nil {
		return nil, err
	}

	canonical := &lemmyResponse.CommunityResponse{
		CommunityView:       converter.ConvertCommunityView(resp.CommunityView),
		DiscussionLanguages: resp.DiscussionLanguages,
	}

	return &http.Response{StatusCode: goHttp.StatusOK, Body: receiver.frontend.BuildCommunityResponse(canonical)}, nil
}

func (receiver *CommunityController) BlockCommunity(request *http.Request) (*http.Response, error) {
	reqDto, err := receiver.frontend.ParseBlockCommunityRequest(request)
	if err != nil {
		return helper.ConvertValidationErrorsToResponse(err), nil
	}

	resp, err := receiver.piefed.BlockCommunity(&piefed.BlockCommunityRequest{
		CommunityId: reqDto.CommunityId,
		Block:       reqDto.Block,
	}, request.Headers)
	if err != nil {
		return nil, err
	}

	canonical := &lemmyResponse.BlockCommunityResponse{
		Blocked:       resp.Blocked,
		CommunityView: converter.ConvertCommunityView(resp.CommunityView),
	}

	return &http.Response{StatusCode: goHttp.StatusOK, Body: receiver.frontend.BuildBlockCommunityResponse(canonical)}, nil
}
