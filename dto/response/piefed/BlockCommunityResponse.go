package piefed

import "LemmyPiefedApi/dto/model/piefed"

type BlockCommunityResponse struct {
	Blocked       bool                  `json:"blocked" required:"true"`
	CommunityView piefed.CommunityView `json:"community_view" required:"true"`
}
