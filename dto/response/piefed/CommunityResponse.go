package piefed

import "LemmyPiefedApi/dto/model/piefed"

type CommunityResponse struct {
	CommunityView       piefed.CommunityView `json:"community_view" required:"true"`
	DiscussionLanguages []uint               `json:"discussion_languages" required:"true"`
}
