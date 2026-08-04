package controller

import (
	"LemmyBeProxy/dto/request/lemmy"
	"LemmyBeProxy/helper"
	"LemmyBeProxy/http"
	"LemmyBeProxy/service/backend"
	goHttp "net/http"
)

// CommentController is thin — all Piefed-specific translation moved into
// PiefedBackend. See PostController for the same pattern applied first.
type CommentController struct {
	backend backend.Backend
}

func NewCommentController(backend backend.Backend) *CommentController {
	return &CommentController{
		backend: backend,
	}
}

func (receiver *CommentController) GetComments(request *http.Request) (*http.Response, error) {
	reqDto, err := helper.ParseRequestQuery[lemmy.GetCommentsRequest](request)
	if err != nil {
		return helper.ConvertValidationErrorsToResponse(err), nil
	}

	resp, err := receiver.backend.GetComments(reqDto, request.Headers)
	if err != nil {
		return nil, err
	}

	return &http.Response{StatusCode: goHttp.StatusOK, Body: resp}, nil
}

func (receiver *CommentController) GetComment(request *http.Request) (*http.Response, error) {
	reqDto, err := helper.ParseRequestQuery[lemmy.GetCommentRequest](request)
	if err != nil {
		return helper.ConvertValidationErrorsToResponse(err), nil
	}

	resp, err := receiver.backend.GetComment(reqDto, request.Headers)
	if err != nil {
		return nil, err
	}

	return &http.Response{StatusCode: goHttp.StatusOK, Body: resp}, nil
}

func (receiver *CommentController) CreateComment(request *http.Request) (*http.Response, error) {
	reqDto, err := helper.ParseRequest[lemmy.CreateCommentRequest](request)
	if err != nil {
		return helper.ConvertValidationErrorsToResponse(err), nil
	}

	resp, err := receiver.backend.CreateComment(reqDto, request.Headers)
	if err != nil {
		return nil, err
	}

	return &http.Response{StatusCode: goHttp.StatusOK, Body: resp}, nil
}

func (receiver *CommentController) LikeComment(request *http.Request) (*http.Response, error) {
	reqDto, err := helper.ParseRequest[lemmy.CreateCommentLikeRequest](request)
	if err != nil {
		return helper.ConvertValidationErrorsToResponse(err), nil
	}

	resp, err := receiver.backend.LikeComment(reqDto, request.Headers)
	if err != nil {
		return nil, err
	}

	return &http.Response{StatusCode: goHttp.StatusOK, Body: resp}, nil
}
