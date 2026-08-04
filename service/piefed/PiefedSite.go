package piefed

import (
	piefedResponse "LemmyBeProxy/dto/response/piefed"
	"LemmyBeProxy/http"
	"LemmyBeProxy/router"
)

func (receiver *Piefed) Site(headers http.Headers) (*piefedResponse.GetSiteResponse, error) {
	return defaultHandler[piefedResponse.GetSiteResponse](
		receiver,
		"/site",
		router.HttpMethodGet,
		nil,
		headers,
	)
}
