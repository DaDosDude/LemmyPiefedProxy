package lemmy

import "LemmyPiefedApi/dto/model/lemmy"

type GetCommunityResponse struct {
	CommunityView lemmy.CommunityView            `json:"community_view" required:"true"`
	Moderators    []lemmy.CommunityModeratorView `json:"moderators" required:"true"`
}
