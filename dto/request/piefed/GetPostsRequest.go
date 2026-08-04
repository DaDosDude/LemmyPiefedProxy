package piefed

import "LemmyBeProxy/dto/model/piefed"

type GetPostsRequest struct {
	Type          *piefed.ListingType `json:"type_,omitempty"`
	Sort          *piefed.SortType    `json:"sort,omitempty"`
	Page          *uint               `json:"page,omitempty"`
	Limit         *uint               `json:"limit,omitempty"`
	CommunityId   *uint               `json:"community_id,omitempty"`
	PersonId      *uint               `json:"person_id,omitempty"`
	CommunityName *string             `json:"community_name,omitempty"`
	LikedOnly     *bool               `json:"liked_only,omitempty"`
}
