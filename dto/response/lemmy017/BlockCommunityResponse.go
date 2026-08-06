package lemmy017

import model017 "LemmyBeProxy/dto/model/lemmy017"

type BlockCommunityResponse struct {
	Blocked       bool                   `json:"blocked" validate:"required"`
	CommunityView model017.CommunityView `json:"community_view" validate:"required"`
}
