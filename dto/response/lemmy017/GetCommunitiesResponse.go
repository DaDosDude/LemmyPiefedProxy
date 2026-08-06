package lemmy017

import model017 "LemmyBeProxy/dto/model/lemmy017"

type GetCommunitiesResponse struct {
	Communities []model017.CommunityView `json:"communities" validate:"required"`
}
