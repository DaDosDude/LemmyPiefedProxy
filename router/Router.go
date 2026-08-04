package router

import (
	"LemmyBeProxy/helper"
	"LemmyBeProxy/http"
	"fmt"
	"slices"
	"strings"
)

type HttpMethod string

const (
	HttpMethodAll    HttpMethod = "*"
	HttpMethodGet    HttpMethod = "GET"
	HttpMethodPost   HttpMethod = "POST"
	HttpMethodPut    HttpMethod = "PUT"
	HttpMethodDelete HttpMethod = "DELETE"
	HttpMethodPatch  HttpMethod = "PATCH"
)

func HttpMethodFromString(method string) (HttpMethod, error) {
	allowed := helper.MapSlice(
		[]HttpMethod{HttpMethodGet, HttpMethodPost, HttpMethodPut, HttpMethodDelete, HttpMethodPatch},
		func(in HttpMethod) string {
			return method
		},
	)
	if !slices.Contains(allowed, strings.ToUpper(method)) {
		return HttpMethodAll, fmt.Errorf("invalid http method: %s", method)
	}

	return HttpMethod(strings.ToUpper(method)), nil
}

type ControllerMethod func(request *http.Request) (*http.Response, error)

type Route struct {
	Path             string
	HttpMethod       HttpMethod
	ControllerMethod ControllerMethod
}

func NewRoute(path string, httpMethod HttpMethod, controller ControllerMethod) *Route {
	return &Route{Path: path, ControllerMethod: controller, HttpMethod: httpMethod}
}

type Router struct {
	Routes []*Route
}

func NewRouter() *Router {
	return &Router{
		Routes: make([]*Route, 0),
	}
}

func (receiver *Router) AddRoute(route *Route) {
	receiver.Routes = append(receiver.Routes, route)
}
