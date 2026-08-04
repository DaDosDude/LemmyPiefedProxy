package piefed

import (
	"LemmyBeProxy/dto/request/piefed"
	piefedResponse "LemmyBeProxy/dto/response/piefed"
	"LemmyBeProxy/http"
	"LemmyBeProxy/router"
)

func (receiver *Piefed) Search(request *piefed.SearchRequest, headers http.Headers) (*piefedResponse.SearchResponse, error) {
	return defaultHandler[piefedResponse.SearchResponse](
		receiver,
		"/search",
		router.HttpMethodGet,
		request,
		headers,
	)
}
