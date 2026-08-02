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

type UserController struct {
	piefed *pfService.Piefed
}

func NewUserController(piefed *pfService.Piefed) *UserController {
	return &UserController{
		piefed: piefed,
	}
}

func (receiver *UserController) Login(request *http.Request) (*http.Response, error) {
	reqDto, err := helper.ParseRequest[lemmy.LoginRequest](request)
	if err != nil {
		return helper.ConvertValidationErrorsToResponse(err), nil
	}

	resp, err := receiver.piefed.Login(&piefed.LoginRequest{
		Username: reqDto.UsernameOrEmail,
		Password: reqDto.Password,
	}, request.Headers)
	if err != nil {
		return nil, err
	}

	return &http.Response{
		StatusCode: goHttp.StatusOK,
		Body: &lemmyResponse.LoginResponse{
			Jwt: resp.Jwt,
		},
	}, nil
}

func (receiver *UserController) Register(request *http.Request) (*http.Response, error) {
	_, err := helper.ParseRequest[lemmy.RegisterRequest](request)
	if err != nil {
		return helper.ConvertValidationErrorsToResponse(err), nil
	}

	return http.NotImplementedResponse(), nil
}

func (receiver *UserController) GetUnreadCount(request *http.Request) (*http.Response, error) {
	resp, err := receiver.piefed.GetUnreadCount(request.Headers)
	if err != nil {
		return nil, err
	}

	return &http.Response{
		StatusCode: goHttp.StatusOK,
		Body: &lemmyResponse.GetUnreadCountResponse{
			Mentions:        resp.Mentions,
			PrivateMessages: resp.PrivateMessages,
			Replies:         resp.Replies,
		},
	}, nil
}

func (receiver *UserController) GetReportCount(request *http.Request) (*http.Response, error) {
	reqDto, err := helper.ParseRequestQuery[lemmy.GetReportCountRequest](request)
	if err != nil {
		return helper.ConvertValidationErrorsToResponse(err), nil
	}

	return &http.Response{
		StatusCode: goHttp.StatusOK,
		Body: &lemmyResponse.GetReportCountResponse{
			CommentReports:        0,
			CommunityId:           reqDto.CommunityId,
			PostReports:           0,
			PrivateMessageReports: nil,
		},
	}, nil
}

func (receiver *UserController) GetUser(request *http.Request) (*http.Response, error) {
	reqDto, err := helper.ParseRequestQuery[lemmy.GetUserRequest](request)
	if err != nil {
		return helper.ConvertValidationErrorsToResponse(err), nil
	}

	resp, err := receiver.piefed.GetUser(&piefed.GetUserRequest{
		PersonId:  reqDto.PersonId,
		Username:  reqDto.Username,
		Sort: helper.SafeDereference(reqDto.Sort, func(in lemmyModel.SortType) *piefedModel.SortType {
			return helper.ToPointer(converter.ReverseConvertSortType(in))
		}),
		Page:      reqDto.Page,
		Limit:     reqDto.Limit,
		SavedOnly: reqDto.SavedOnly,
	}, request.Headers)
	if err != nil {
		return nil, err
	}

	return &http.Response{
		StatusCode: goHttp.StatusOK,
		Body: &lemmyResponse.GetUserResponse{
			Comments:   helper.MapSlice(resp.Comments, converter.ConvertCommentView),
			Moderates:  helper.MapSlice(resp.Moderates, converter.ConvertCommunityModeratorView),
			PersonView: converter.ConvertPersonView(resp.PersonView),
			Posts:      helper.MapSlice(resp.Posts, converter.ConvertPostView),
		},
	}, nil
}

func (receiver *UserController) BlockPerson(request *http.Request) (*http.Response, error) {
	reqDto, err := helper.ParseRequest[lemmy.BlockPersonRequest](request)
	if err != nil {
		return helper.ConvertValidationErrorsToResponse(err), nil
	}

	resp, err := receiver.piefed.BlockPerson(&piefed.BlockPersonRequest{
		PersonId: reqDto.PersonId,
		Block:    reqDto.Block,
	}, request.Headers)
	if err != nil {
		return nil, err
	}

	return &http.Response{
		StatusCode: goHttp.StatusOK,
		Body: &lemmyResponse.BlockPersonResponse{
			Blocked:    resp.Blocked,
			PersonView: converter.ConvertPersonView(resp.PersonView),
		},
	}, nil
}
