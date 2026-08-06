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
	"LemmyBeProxy/service/frontend"
	pfService "LemmyBeProxy/service/piefed"
	goHttp "net/http"
)

// UserController uses frontend.Frontend for the endpoints migrated so
// far (Login, GetUnreadCount, GetUser, BlockPerson). Register,
// GetReportCount, and SaveUserSettings are unmigrated — Register and
// GetReportCount are stub responses regardless of wire format, and
// SaveUserSettings's 0.17.x translation is deferred (numeric sort-type
// enum indices, a genuinely separate problem — see the project's
// Features not working yet).
type UserController struct {
	piefed   *pfService.Piefed
	frontend frontend.Frontend
}

func NewUserController(piefed *pfService.Piefed, frontend frontend.Frontend) *UserController {
	return &UserController{
		piefed:   piefed,
		frontend: frontend,
	}
}

func (receiver *UserController) Login(request *http.Request) (*http.Response, error) {
	reqDto, err := receiver.frontend.ParseLoginRequest(request)
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

	canonical := &lemmyResponse.LoginResponse{
		Jwt: resp.Jwt,
	}

	return &http.Response{StatusCode: goHttp.StatusOK, Body: receiver.frontend.BuildLoginResponse(canonical)}, nil
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

	canonical := &lemmyResponse.GetUnreadCountResponse{
		Mentions:        resp.Mentions,
		PrivateMessages: resp.PrivateMessages,
		Replies:         resp.Replies,
	}

	return &http.Response{StatusCode: goHttp.StatusOK, Body: receiver.frontend.BuildGetUnreadCountResponse(canonical)}, nil
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
	reqDto, err := receiver.frontend.ParseGetUserRequest(request)
	if err != nil {
		return helper.ConvertValidationErrorsToResponse(err), nil
	}

	resp, err := receiver.piefed.GetUser(&piefed.GetUserRequest{
		PersonId: reqDto.PersonId,
		Username: reqDto.Username,
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

	canonical := &lemmyResponse.GetUserResponse{
		Comments:   helper.MapSlice(resp.Comments, converter.ConvertCommentView),
		Moderates:  helper.MapSlice(resp.Moderates, converter.ConvertCommunityModeratorView),
		PersonView: converter.ConvertPersonView(resp.PersonView),
		Posts:      helper.MapSlice(resp.Posts, converter.ConvertPostView),
	}

	return &http.Response{StatusCode: goHttp.StatusOK, Body: receiver.frontend.BuildGetUserResponse(canonical)}, nil
}

func (receiver *UserController) BlockPerson(request *http.Request) (*http.Response, error) {
	reqDto, err := receiver.frontend.ParseBlockPersonRequest(request)
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

	canonical := &lemmyResponse.BlockPersonResponse{
		Blocked:    resp.Blocked,
		PersonView: converter.ConvertPersonView(resp.PersonView),
	}

	return &http.Response{StatusCode: goHttp.StatusOK, Body: receiver.frontend.BuildBlockPersonResponse(canonical)}, nil
}

// SaveUserSettings only forwards the four fields Piefed's own
// save_user_settings endpoint actually supports (ShowNsfw, DefaultSortType,
// DefaultCommentSortType, ShowReadPosts). Everything else mlmym sends
// (Theme, InfiniteScrollEnabled, ShowAvatars, etc.) is accepted here — so
// the save doesn't fail outright — but has no Piefed field to go to and is
// silently dropped. InfiniteScrollEnabled specifically can never persist
// server-side on a Piefed-backed instance regardless of what this proxy
// does, since Piefed has no field for it at all.
//
// Not migrated onto frontend.Frontend yet: 0.17.x's real SaveUserSettings
// encodes default_sort_type/default_listing_type as raw numeric enum
// indices rather than string names, unlike every other sort-bearing
// endpoint — a genuinely separate translation problem, deferred rather
// than rushed.
func (receiver *UserController) SaveUserSettings(request *http.Request) (*http.Response, error) {
	reqDto, err := helper.ParseRequest[lemmy.SaveUserSettingsRequest](request)
	if err != nil {
		return helper.ConvertValidationErrorsToResponse(err), nil
	}

	_, err = receiver.piefed.SaveUserSettings(&piefed.SaveUserSettingsRequest{
		ShowNsfw: reqDto.ShowNsfw,
		DefaultSortType: helper.SafeDereference(reqDto.DefaultSortType, func(in string) *string {
			return helper.ToPointer(converter.ClampDefaultSortType(in))
		}),
		DefaultCommentSortType: helper.SafeDereference(reqDto.DefaultCommentSortType, func(in string) *string {
			return helper.ToPointer(converter.ClampDefaultCommentSortType(in))
		}),
		ShowReadPosts: reqDto.ShowReadPosts,
	}, request.Headers)
	if err != nil {
		return nil, err
	}

	return &http.Response{
		StatusCode: goHttp.StatusOK,
		Body:       &lemmyResponse.SaveUserSettingsResponse{},
	}, nil
}
