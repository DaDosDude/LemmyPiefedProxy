package lemmy

import "LemmyPiefedApi/dto/model/lemmy"

type BlockCommunityResponse struct {
	Blocked       bool                `json:"blocked" required:"true"`
	CommunityView lemmy.CommunityView `json:"community_view" required:"true"`
}
