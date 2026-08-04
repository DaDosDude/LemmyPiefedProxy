package piefed

import (
	"LemmyBeProxy/dto/request/piefed"
	piefedResponse "LemmyBeProxy/dto/response/piefed"
	"LemmyBeProxy/http"
	"LemmyBeProxy/router"
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

func (receiver *Piefed) BlockCommunity(request *piefed.BlockCommunityRequest, headers http.Headers) (*piefedResponse.BlockCommunityResponse, error) {
	return defaultHandler[piefedResponse.BlockCommunityResponse](
		receiver,
		"/community/block",
		router.HttpMethodPost,
		request,
		headers,
	)
}
