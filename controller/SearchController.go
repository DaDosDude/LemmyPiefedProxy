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

type SearchController struct {
	piefed   *pfService.Piefed
	frontend frontend.Frontend
}

func NewSearchController(piefed *pfService.Piefed, frontend frontend.Frontend) *SearchController {
	return &SearchController{
		piefed:   piefed,
		frontend: frontend,
	}
}

func (receiver *SearchController) Search(request *http.Request) (*http.Response, error) {
	reqDto, err := receiver.frontend.ParseSearchRequest(request)
	if err != nil {
		return helper.ConvertValidationErrorsToResponse(err), nil
	}

	resp, err := receiver.piefed.Search(&piefed.SearchRequest{
		Q:    reqDto.Q,
		Type: converter.ReverseConvertSearchType(reqDto.Type),
		ListingType: helper.SafeDereference(reqDto.ListingType, func(in lemmyModel.ListingType) *piefedModel.ListingType {
			return helper.ToPointer(converter.ReverseConvertListingType(in))
		}),
		Sort: helper.SafeDereference(reqDto.Sort, func(in lemmyModel.SortType) *piefedModel.SortType {
			return helper.ToPointer(converter.ReverseConvertSortType(in))
		}),
		Page:          reqDto.Page,
		Limit:         reqDto.Limit,
		CommunityName: reqDto.CommunityName,
		CommunityId:   reqDto.CommunityId,
	}, request.Headers)
	if err != nil {
		return nil, err
	}

	canonical := &lemmyResponse.SearchResponse{
		Type:        converter.ConvertSearchType(resp.Type),
		Communities: helper.MapSlice(resp.Communities, converter.ConvertCommunityView),
		Posts:       helper.MapSlice(resp.Posts, converter.ConvertPostView),
		Users:       helper.MapSlice(resp.Users, converter.ConvertPersonView),
		Comments:    helper.MapSlice(resp.Comments, converter.ConvertCommentView),
	}

	return &http.Response{StatusCode: goHttp.StatusOK, Body: receiver.frontend.BuildSearchResponse(canonical)}, nil
}
