package piefed

import (
	"LemmyPiefedApi/dto/request/piefed"
	piefedResponse "LemmyPiefedApi/dto/response/piefed"
	"LemmyPiefedApi/http"
	"LemmyPiefedApi/router"
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
