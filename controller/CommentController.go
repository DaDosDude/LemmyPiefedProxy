package controller

import (
	"LemmyBeProxy/helper"
	"LemmyBeProxy/http"
	"LemmyBeProxy/service/backend"
	"LemmyBeProxy/service/frontend"
	goHttp "net/http"
)

// CommentController is thin on both ends, same pattern as PostController.
type CommentController struct {
	backend  backend.Backend
	frontend frontend.Frontend
}

func NewCommentController(backend backend.Backend, frontend frontend.Frontend) *CommentController {
	return &CommentController{
		backend:  backend,
		frontend: frontend,
	}
}

func (receiver *CommentController) GetComments(request *http.Request) (*http.Response, error) {
	reqDto, err := receiver.frontend.ParseGetCommentsRequest(request)
	if err != nil {
		return helper.ConvertValidationErrorsToResponse(err), nil
	}

	resp, err := receiver.backend.GetComments(reqDto, request.Headers)
	if err != nil {
		return nil, err
	}

	return &http.Response{StatusCode: goHttp.StatusOK, Body: receiver.frontend.BuildGetCommentsResponse(resp)}, nil
}

func (receiver *CommentController) GetComment(request *http.Request) (*http.Response, error) {
	reqDto, err := receiver.frontend.ParseGetCommentRequest(request)
	if err != nil {
		return helper.ConvertValidationErrorsToResponse(err), nil
	}

	resp, err := receiver.backend.GetComment(reqDto, request.Headers)
	if err != nil {
		return nil, err
	}

	return &http.Response{StatusCode: goHttp.StatusOK, Body: receiver.frontend.BuildGetCommentResponse(resp)}, nil
}

func (receiver *CommentController) CreateComment(request *http.Request) (*http.Response, error) {
	reqDto, err := receiver.frontend.ParseCreateCommentRequest(request)
	if err != nil {
		return helper.ConvertValidationErrorsToResponse(err), nil
	}

	resp, err := receiver.backend.CreateComment(reqDto, request.Headers)
	if err != nil {
		return nil, err
	}

	return &http.Response{StatusCode: goHttp.StatusOK, Body: receiver.frontend.BuildCreateCommentResponse(resp)}, nil
}

func (receiver *CommentController) LikeComment(request *http.Request) (*http.Response, error) {
	reqDto, err := receiver.frontend.ParseCreateCommentLikeRequest(request)
	if err != nil {
		return helper.ConvertValidationErrorsToResponse(err), nil
	}

	resp, err := receiver.backend.LikeComment(reqDto, request.Headers)
	if err != nil {
		return nil, err
	}

	return &http.Response{StatusCode: goHttp.StatusOK, Body: receiver.frontend.BuildLikeCommentResponse(resp)}, nil
}
