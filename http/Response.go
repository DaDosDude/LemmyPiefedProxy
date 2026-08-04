package http

import (
	"LemmyBeProxy/dto/response"
	"net/http"
)

type Response struct {
	StatusCode int
	Body       any
	Headers    map[string]string
}

func InternalProxyError() *Response {
	return &Response{
		StatusCode: http.StatusInternalServerError,
		Body: &response.InternalError{
			Error: "internal lemmy layer proxy error",
		},
	}
}

func NotFoundProxyError() *Response {
	return &Response{
		StatusCode: http.StatusNotFound,
		Body: &response.InternalError{
			Error: "not found",
		},
	}
}

func NotImplementedResponse() *Response {
	return &Response{
		StatusCode: http.StatusNotImplemented,
		Body: &response.InternalError{
			Error: "this method has not been implemented yet in the proxy (or piefed)",
		},
	}
}

func NotImplementedFeature(description string) *Response {
	return &Response{
		StatusCode: http.StatusNotImplemented,
		Body: &response.InternalError{
			Error: description,
		},
	}
}

func NoContent() *Response {
	return &Response{
		StatusCode: http.StatusNoContent,
	}
}
