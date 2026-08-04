package controller

import (
	lemmyModel "LemmyBeProxy/dto/model/lemmy"
	piefedModel "LemmyBeProxy/dto/model/piefed"
	"LemmyBeProxy/dto/request/lemmy"
	"LemmyBeProxy/dto/request/piefed"
	lemmyResponse "LemmyBeProxy/dto/response/lemmy"
	"LemmyBeProxy/helper"
	"LemmyBeProxy/helper/converter"
	"LemmyBeProxy/http"
	pfService "LemmyBeProxy/service/piefed"
	goHttp "net/http"
)

type CommunityController struct {
	piefed *pfService.Piefed
}

func NewCommunityController(piefed *pfService.Piefed) *CommunityController {
	return &CommunityController{
		piefed: piefed,
	}
}

func (receiver *CommunityController) GetCommunity(request *http.Request) (*http.Response, error) {
	reqDto, err := helper.ParseRequestQuery[lemmy.GetCommunityRequest](request)
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

	return &http.Response{
		StatusCode: goHttp.StatusOK,
		Body: &lemmyResponse.GetCommunityResponse{
			CommunityView: converter.ConvertCommunityView(resp.CommunityView),
			Moderators:    helper.MapSlice(resp.Moderators, converter.ConvertCommunityModeratorView),
		},
	}, nil
}

func (receiver *CommunityController) GetCommunities(request *http.Request) (*http.Response, error) {
	reqDto, err := helper.ParseRequestQuery[lemmy.GetCommunitiesRequest](request)
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

	return &http.Response{
		StatusCode: goHttp.StatusOK,
		Body: &lemmyResponse.GetCommunitiesResponse{
			Communities: helper.MapSlice(resp.Communities, converter.ConvertCommunityView),
		},
	}, nil
}

func (receiver *CommunityController) FollowCommunity(request *http.Request) (*http.Response, error) {
	reqDto, err := helper.ParseRequest[lemmy.FollowCommunityRequest](request)
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

	return &http.Response{
		StatusCode: goHttp.StatusOK,
		Body: &lemmyResponse.CommunityResponse{
			CommunityView:       converter.ConvertCommunityView(resp.CommunityView),
			DiscussionLanguages: resp.DiscussionLanguages,
		},
	}, nil
}

func (receiver *CommunityController) BlockCommunity(request *http.Request) (*http.Response, error) {
	reqDto, err := helper.ParseRequest[lemmy.BlockCommunityRequest](request)
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

	return &http.Response{
		StatusCode: goHttp.StatusOK,
		Body: &lemmyResponse.BlockCommunityResponse{
			Blocked:       resp.Blocked,
			CommunityView: converter.ConvertCommunityView(resp.CommunityView),
		},
	}, nil
}
