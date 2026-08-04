package controller

import (
	"LemmyBeProxy/dto/request/lemmy"
	"LemmyBeProxy/helper"
	"LemmyBeProxy/http"
	"LemmyBeProxy/service/backend"
	goHttp "net/http"
)

// PostController is now thin — it only parses HTTP in, calls the
// backend, and wraps the result back into an HTTP response. All the
// Piefed-specific field translation that used to live here (sort/listing
// type conversion, the community_id + Subscribed workaround, mapping
// Lemmy's Name to Piefed's Title, etc.) moved into PiefedBackend, since
// that's backend-specific logic a Lemmy backend doesn't need at all.
type PostController struct {
	backend backend.Backend
}

func NewPostController(backend backend.Backend) *PostController {
	return &PostController{
		backend: backend,
	}
}

func (receiver *PostController) GetPosts(request *http.Request) (*http.Response, error) {
	reqDto, err := helper.ParseRequestQuery[lemmy.GetPostsRequest](request)
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

	return &http.Response{StatusCode: goHttp.StatusOK, Body: resp}, nil
}

func (receiver *PostController) GetPost(request *http.Request) (*http.Response, error) {
	reqDto, err := helper.ParseRequestQuery[lemmy.GetPostRequest](request)
	if err != nil {
		return helper.ConvertValidationErrorsToResponse(err), nil
	}

	resp, err := receiver.backend.GetPost(reqDto, request.Headers)
	if err != nil {
		return nil, err
	}

	return &http.Response{StatusCode: goHttp.StatusOK, Body: resp}, nil
}

func (receiver *PostController) MarkPostAsRead(request *http.Request) (*http.Response, error) {
	reqDto, err := helper.ParseRequest[lemmy.MarkPostAsReadRequest](request)
	if err != nil {
		return helper.ConvertValidationErrorsToResponse(err), nil
	}

	resp, err := receiver.backend.MarkPostAsRead(reqDto, request.Headers)
	if err != nil {
		return nil, err
	}

	return &http.Response{StatusCode: goHttp.StatusOK, Body: resp}, nil
}

func (receiver *PostController) CreatePost(request *http.Request) (*http.Response, error) {
	reqDto, err := helper.ParseRequest[lemmy.CreatePostRequest](request)
	if err != nil {
		return helper.ConvertValidationErrorsToResponse(err), nil
	}

	resp, err := receiver.backend.CreatePost(reqDto, request.Headers)
	if err != nil {
		return nil, err
	}

	return &http.Response{StatusCode: goHttp.StatusOK, Body: resp}, nil
}

func (receiver *PostController) EditPost(request *http.Request) (*http.Response, error) {
	reqDto, err := helper.ParseRequest[lemmy.EditPostRequest](request)
	if err != nil {
		return helper.ConvertValidationErrorsToResponse(err), nil
	}

	resp, err := receiver.backend.EditPost(reqDto, request.Headers)
	if err != nil {
		return nil, err
	}

	return &http.Response{StatusCode: goHttp.StatusOK, Body: resp}, nil
}

func (receiver *PostController) LikePost(request *http.Request) (*http.Response, error) {
	reqDto, err := helper.ParseRequest[lemmy.CreatePostLikeRequest](request)
	if err != nil {
		return helper.ConvertValidationErrorsToResponse(err), nil
	}

	resp, err := receiver.backend.LikePost(reqDto, request.Headers)
	if err != nil {
		return nil, err
	}

	return &http.Response{StatusCode: goHttp.StatusOK, Body: resp}, nil
}
