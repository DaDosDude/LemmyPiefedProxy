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
	piefedService "LemmyPiefedApi/service/piefed"
	goHttp "net/http"
)

type CommentController struct {
	piefed *piefedService.Piefed
}

func NewCommentController(piefed *piefedService.Piefed) *CommentController {
	return &CommentController{
		piefed: piefed,
	}
}

func (receiver *CommentController) GetComments(request *http.Request) (*http.Response, error) {
	reqDto, err := helper.ParseRequestQuery[lemmy.GetCommentsRequest](request)
	if err != nil {
		return helper.ConvertValidationErrorsToResponse(err), nil
	}

	resp, err := receiver.piefed.GetComments(&piefed.GetCommentsRequest{
		Type: helper.SafeDereference(reqDto.Type, func(in lemmyModel.ListingType) *piefedModel.ListingType {
			return helper.ToPointer(converter.ReverseConvertListingType(in))
		}),
		PersonId:    nil,
		MaxDepth:    reqDto.MaxDepth,
		Page:        reqDto.Page,
		ParentId:    reqDto.ParentId,
		CommunityId: reqDto.CommunityId,
		PostId:      reqDto.PostId,
		Limit:       reqDto.Limit,
		Sort: helper.SafeDereference(reqDto.Sort, func(in lemmyModel.CommentSortType) *piefedModel.CommentSortType {
			return helper.ToPointer(converter.ReverseConvertCommentSortType(in))
		}),
	}, request.Headers)
	if err != nil {
		return nil, err
	}

	return &http.Response{
		StatusCode: goHttp.StatusOK,
		Body: &lemmyResponse.GetCommentsResponse{
			Comments: helper.MapSlice(resp.Comments, converter.ConvertCommentView),
		},
	}, nil
}

func (receiver *CommentController) GetComment(request *http.Request) (*http.Response, error) {
	reqDto, err := helper.ParseRequestQuery[lemmy.GetCommentRequest](request)
	if err != nil {
		return helper.ConvertValidationErrorsToResponse(err), nil
	}

	resp, err := receiver.piefed.GetComment(&piefed.GetCommentRequest{
		Id: reqDto.Id,
	}, request.Headers)
	if err != nil {
		return nil, err
	}

	return &http.Response{
		StatusCode: goHttp.StatusOK,
		Body: &lemmyResponse.GetCommentResponse{
			CommentView:  converter.ConvertCommentView(resp.CommentView),
			RecipientIds: []uint{},
		},
	}, nil
}

func (receiver *CommentController) CreateComment(request *http.Request) (*http.Response, error) {
	reqDto, err := helper.ParseRequest[lemmy.CreateCommentRequest](request)
	if err != nil {
		return helper.ConvertValidationErrorsToResponse(err), nil
	}

	resp, err := receiver.piefed.CreateComment(&piefed.CreateCommentRequest{
		Body:       reqDto.Content,
		PostId:     reqDto.PostId,
		ParentId:   reqDto.ParentId,
		LanguageId: reqDto.LanguageId,
	}, request.Headers)
	if err != nil {
		return nil, err
	}

	return &http.Response{
		StatusCode: goHttp.StatusOK,
		Body: &lemmyResponse.CreateCommentResponse{
			CommentView:  converter.ConvertCommentView(resp.CommentView),
			RecipientIds: []uint{},
		},
	}, nil
}

func (receiver *CommentController) LikeComment(request *http.Request) (*http.Response, error) {
	reqDto, err := helper.ParseRequest[lemmy.CreateCommentLikeRequest](request)
	if err != nil {
		return helper.ConvertValidationErrorsToResponse(err), nil
	}

	resp, err := receiver.piefed.LikeComment(&piefed.LikeCommentRequest{
		CommentId: reqDto.CommentId,
		Score:     reqDto.Score,
	}, request.Headers)
	if err != nil {
		return nil, err
	}

	return &http.Response{
		StatusCode: goHttp.StatusOK,
		Body: &lemmyResponse.GetCommentResponse{
			CommentView:  converter.ConvertCommentView(resp.CommentView),
			RecipientIds: []uint{},
		},
	}, nil
}
