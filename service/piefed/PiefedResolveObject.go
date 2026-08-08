package piefed

import (
	"LemmyBeProxy/dto/request/piefed"
	piefedResponse "LemmyBeProxy/dto/response/piefed"
	"LemmyBeProxy/http"
	"LemmyBeProxy/router"
)

func (receiver *Piefed) ResolveObject(request *piefed.ResolveObjectRequest, headers http.Headers) (*piefedResponse.ResolveObjectResponse, error) {
	return defaultHandler[piefedResponse.ResolveObjectResponse](
		receiver,
		"/resolve_object",
		router.HttpMethodGet,
		request,
		headers,
	)
}
