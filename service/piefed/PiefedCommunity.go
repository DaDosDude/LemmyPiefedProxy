package piefed

import (
	"LemmyPiefedApi/dto/request/piefed"
	piefedResponse "LemmyPiefedApi/dto/response/piefed"
	"LemmyPiefedApi/http"
	"LemmyPiefedApi/router"
)

func (receiver *Piefed) GetCommunity(request *piefed.GetCommunityRequest, headers http.Headers) (*piefedResponse.GetCommunityResponse, error) {
	return defaultHandler[piefedResponse.GetCommunityResponse](
		receiver,
		"/community",
		router.HttpMethodGet,
		request,
		headers,
	)
}

func (receiver *Piefed) GetCommunities(request *piefed.GetCommunitiesRequest, headers http.Headers) (*piefedResponse.GetCommunitiesResponse, error) {
	return defaultHandler[piefedResponse.GetCommunitiesResponse](
		receiver,
		"/community/list",
		router.HttpMethodGet,
		request,
		headers,
	)
}

func (receiver *Piefed) FollowCommunity(request *piefed.FollowCommunityRequest, headers http.Headers) (*piefedResponse.CommunityResponse, error) {
	return defaultHandler[piefedResponse.CommunityResponse](
		receiver,
		"/community/follow",
		router.HttpMethodPost,
		request,
		headers,
	)
}
