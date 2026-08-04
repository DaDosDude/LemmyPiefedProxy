package controller

import (
	"LemmyBeProxy/helper"
	"LemmyBeProxy/http"
	"LemmyBeProxy/service/backend"
	"LemmyBeProxy/service/frontend"
	goHttp "net/http"
)

// PostController is thin on both ends now — frontend handles parsing the
// request and building the response in whatever wire format the
// configured FRONTEND_VERSION expects, backend handles translating to
// and from whatever BACKEND_TYPE is actually running. The controller
// itself has no version-specific knowledge of either side at all.
type PostController struct {
	backend  backend.Backend
	frontend frontend.Frontend
}

func NewPostController(backend backend.Backend, frontend frontend.Frontend) *PostController {
	return &PostController{
		backend:  backend,
		frontend: frontend,
	}
}

func (receiver *PostController) GetPosts(request *http.Request) (*http.Response, error) {
	reqDto, err := receiver.frontend.ParseGetPostsRequest(request)
	if err != nil {
		return helper.ConvertValidationErrorsToResponse(err), nil
	}

	if reqDto.PageCursor != nil {
		return http.NotImplementedFeature(
			"Cursor based navigation is not available",
		), nil
	}

	resp, err := receiver.backend.GetPosts(reqDto, request.Headers)
	if err != nil {
		return nil, err
	}

	return &http.Response{StatusCode: goHttp.StatusOK, Body: receiver.frontend.BuildGetPostsResponse(resp)}, nil
}

func (receiver *PostController) GetPost(request *http.Request) (*http.Response, error) {
	reqDto, err := receiver.frontend.ParseGetPostRequest(request)
	if err != nil {
		return helper.ConvertValidationErrorsToResponse(err), nil
	}

	resp, err := receiver.backend.GetPost(reqDto, request.Headers)
	if err != nil {
		return nil, err
	}

	return &http.Response{StatusCode: goHttp.StatusOK, Body: receiver.frontend.BuildGetPostResponse(resp)}, nil
}

func (receiver *PostController) MarkPostAsRead(request *http.Request) (*http.Response, error) {
	reqDto, err := receiver.frontend.ParseMarkPostAsReadRequest(request)
	if err != nil {
		return helper.ConvertValidationErrorsToResponse(err), nil
	}

	resp, err := receiver.backend.MarkPostAsRead(reqDto, request.Headers)
	if err != nil {
		return nil, err
	}

	return &http.Response{StatusCode: goHttp.StatusOK, Body: receiver.frontend.BuildSuccessResponse(resp)}, nil
}

func (receiver *PostController) CreatePost(request *http.Request) (*http.Response, error) {
	reqDto, err := receiver.frontend.ParseCreatePostRequest(request)
	if err != nil {
		return helper.ConvertValidationErrorsToResponse(err), nil
	}

	resp, err := receiver.backend.CreatePost(reqDto, request.Headers)
	if err != nil {
		return nil, err
	}

	return &http.Response{StatusCode: goHttp.StatusOK, Body: receiver.frontend.BuildPostMutationResponse(resp)}, nil
}

func (receiver *PostController) EditPost(request *http.Request) (*http.Response, error) {
	reqDto, err := receiver.frontend.ParseEditPostRequest(request)
	if err != nil {
		return helper.ConvertValidationErrorsToResponse(err), nil
	}

	resp, err := receiver.backend.EditPost(reqDto, request.Headers)
	if err != nil {
		return nil, err
	}

	return &http.Response{StatusCode: goHttp.StatusOK, Body: receiver.frontend.BuildPostMutationResponse(resp)}, nil
}

func (receiver *PostController) LikePost(request *http.Request) (*http.Response, error) {
	reqDto, err := receiver.frontend.ParseCreatePostLikeRequest(request)
	if err != nil {
		return helper.ConvertValidationErrorsToResponse(err), nil
	}

	resp, err := receiver.backend.LikePost(reqDto, request.Headers)
	if err != nil {
		return nil, err
	}

	return &http.Response{StatusCode: goHttp.StatusOK, Body: receiver.frontend.BuildPostMutationResponse(resp)}, nil
}
