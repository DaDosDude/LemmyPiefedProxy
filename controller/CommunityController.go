package controller

import (
	"LemmyBeProxy/helper"
	"LemmyBeProxy/http"
	"LemmyBeProxy/service/backend"
	"LemmyBeProxy/service/frontend"
	goHttp "net/http"
)

// CommunityController is now thin on both axes — frontend handles wire
// format, backend handles which real service is actually called. All
// Piefed-specific translation that used to live in this controller
// moved into PiefedBackend, matching the pattern PostController and
// CommentController were already migrated onto.
type CommunityController struct {
	backend  backend.Backend
	frontend frontend.Frontend
}

func NewCommunityController(backend backend.Backend, frontend frontend.Frontend) *CommunityController {
	return &CommunityController{
		backend:  backend,
		frontend: frontend,
	}
}

func (receiver *CommunityController) GetCommunity(request *http.Request) (*http.Response, error) {
	reqDto, err := receiver.frontend.ParseGetCommunityRequest(request)
	if err != nil {
		return helper.ConvertValidationErrorsToResponse(err), nil
	}

	resp, err := receiver.backend.GetCommunity(reqDto, request.Headers)
	if err != nil {
		return nil, err
	}

	return &http.Response{StatusCode: goHttp.StatusOK, Body: receiver.frontend.BuildGetCommunityResponse(resp)}, nil
}

func (receiver *CommunityController) GetCommunities(request *http.Request) (*http.Response, error) {
	reqDto, err := receiver.frontend.ParseGetCommunitiesRequest(request)
	if err != nil {
		return helper.ConvertValidationErrorsToResponse(err), nil
	}

	resp, err := receiver.backend.GetCommunities(reqDto, request.Headers)
	if err != nil {
		return nil, err
	}

	return &http.Response{StatusCode: goHttp.StatusOK, Body: receiver.frontend.BuildGetCommunitiesResponse(resp)}, nil
}

func (receiver *CommunityController) FollowCommunity(request *http.Request) (*http.Response, error) {
	reqDto, err := receiver.frontend.ParseFollowCommunityRequest(request)
	if err != nil {
		return helper.ConvertValidationErrorsToResponse(err), nil
	}

	resp, err := receiver.backend.FollowCommunity(reqDto, request.Headers)
	if err != nil {
		return nil, err
	}

	return &http.Response{StatusCode: goHttp.StatusOK, Body: receiver.frontend.BuildCommunityResponse(resp)}, nil
}

func (receiver *CommunityController) BlockCommunity(request *http.Request) (*http.Response, error) {
	reqDto, err := receiver.frontend.ParseBlockCommunityRequest(request)
	if err != nil {
		return helper.ConvertValidationErrorsToResponse(err), nil
	}

	resp, err := receiver.backend.BlockCommunity(reqDto, request.Headers)
	if err != nil {
		return nil, err
	}

	return &http.Response{StatusCode: goHttp.StatusOK, Body: receiver.frontend.BuildBlockCommunityResponse(resp)}, nil
}
