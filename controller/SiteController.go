package controller

import (
	"LemmyBeProxy/helper"
	"LemmyBeProxy/http"
	"LemmyBeProxy/service/backend"
	"LemmyBeProxy/service/frontend"
	goHttp "net/http"
)

// SiteController is now thin — all the Piefed-specific assembly logic
// (including the ActivityPub actor fetch) moved into PiefedBackend,
// since that's genuinely backend-specific work a real Lemmy backend
// doesn't need at all.
type SiteController struct {
	backend  backend.Backend
	frontend frontend.Frontend
}

func NewSiteController(backend backend.Backend, frontend frontend.Frontend) *SiteController {
	return &SiteController{
		backend:  backend,
		frontend: frontend,
	}
}

func (receiver *SiteController) Site(request *http.Request) (*http.Response, error) {
	resp, err := receiver.backend.Site(request.Headers)
	if err != nil {
		return nil, err
	}

	return &http.Response{StatusCode: goHttp.StatusOK, Body: receiver.frontend.BuildGetSiteResponse(resp)}, nil
}

func (receiver *SiteController) ResolveObject(request *http.Request) (*http.Response, error) {
	reqDto, err := receiver.frontend.ParseResolveObjectRequest(request)
	if err != nil {
		return helper.ConvertValidationErrorsToResponse(err), nil
	}

	resp, err := receiver.backend.ResolveObject(reqDto, request.Headers)
	if err != nil {
		return nil, err
	}

	return &http.Response{StatusCode: goHttp.StatusOK, Body: receiver.frontend.BuildResolveObjectResponse(resp)}, nil
}
