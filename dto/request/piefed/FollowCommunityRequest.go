package piefed

type FollowCommunityRequest struct {
	CommunityId uint `json:"community_id" validate:"required"`
	Follow      bool `json:"follow"`
}
