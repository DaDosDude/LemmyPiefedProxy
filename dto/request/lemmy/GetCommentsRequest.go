package lemmy

import "LemmyBeProxy/dto/model/lemmy"

type GetCommentsRequest struct {
	CommunityId   *uint                  `json:"community_id,omitempty"`
	CommunityName *string                `json:"community_name,omitempty"`
	DislikedOnly  *bool                  `json:"disliked_only,omitempty"`
	LikedOnly     *bool                  `json:"liked_only,omitempty"`
	Limit         *uint                  `json:"limit,omitempty"`
	MaxDepth      *uint                  `json:"max_depth,omitempty"`
	Page          *uint                  `json:"page,omitempty"`
	ParentId      *uint                  `json:"parent_id,omitempty"`
	PostId        *uint                  `json:"post_id,omitempty"`
	SavedOnly     *bool                  `json:"saved_only,omitempty"`
	Sort          *lemmy.CommentSortType `json:"sort,omitempty"`
	Type          *lemmy.ListingType     `json:"type_,omitempty"`
}
