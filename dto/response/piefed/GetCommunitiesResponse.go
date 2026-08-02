package piefed

import "LemmyPiefedApi/dto/model/piefed"

type GetCommunitiesResponse struct {
	Communities []piefed.CommunityView `json:"communities" required:"true"`
}
