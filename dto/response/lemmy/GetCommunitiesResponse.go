package lemmy

import "LemmyBeProxy/dto/model/lemmy"

type GetCommunitiesResponse struct {
	Communities []lemmy.CommunityView `json:"communities" required:"true"`
}
