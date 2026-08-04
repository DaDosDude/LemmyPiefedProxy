package piefed

import "LemmyBeProxy/dto/model/piefed"

type GetCommunityResponse struct {
	CommunityView piefed.CommunityView            `json:"community_view" required:"true"`
	Moderators    []piefed.CommunityModeratorView `json:"moderators" required:"true"`
}
