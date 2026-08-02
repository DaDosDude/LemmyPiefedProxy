package lemmy

import "LemmyPiefedApi/dto/model/lemmy"

type GetCommunitiesResponse struct {
	Communities []lemmy.CommunityView `json:"communities" required:"true"`
}
