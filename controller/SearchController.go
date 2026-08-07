package controller

import (
	"LemmyBeProxy/helper"
	"LemmyBeProxy/http"
	"LemmyBeProxy/service/backend"
	"LemmyBeProxy/service/frontend"
	goHttp "net/http"
)

type SearchController struct {
	backend  backend.Backend
	frontend frontend.Frontend
}

func NewSearchController(backend backend.Backend, frontend frontend.Frontend) *SearchController {
	return &SearchController{
		backend:  backend,
		frontend: frontend,
	}
}

func (receiver *SearchController) Search(request *http.Request) (*http.Response, error) {
	reqDto, err := receiver.frontend.ParseSearchRequest(request)
	if err != nil {
		return helper.ConvertValidationErrorsToResponse(err), nil
	}

	resp, err := receiver.backend.Search(reqDto, request.Headers)
	if err != nil {
		return nil, err
	}

	return &http.Response{StatusCode: goHttp.StatusOK, Body: receiver.frontend.BuildSearchResponse(resp)}, nil
}
