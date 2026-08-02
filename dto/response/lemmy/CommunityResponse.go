package lemmy

import "LemmyPiefedApi/dto/model/lemmy"

type CommunityResponse struct {
	CommunityView       lemmy.CommunityView `json:"community_view" required:"true"`
	DiscussionLanguages []uint              `json:"discussion_languages" required:"true"`
}
