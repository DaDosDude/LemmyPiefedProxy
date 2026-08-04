package lemmy017

import "LemmyBeProxy/dto/model/lemmy"

// GetCommentsRequest matches Lemmy 0.17.2's real GetComments struct
// exactly. saved_only exists here too, same as GetPosts.
type GetCommentsRequest struct {
	Type          *lemmy.ListingType     `json:"type_,omitempty"`
	Sort          *lemmy.CommentSortType `json:"sort,omitempty"`
	MaxDepth      *uint                  `json:"max_depth,omitempty"`
	Page          *uint                  `json:"page,omitempty"`
	Limit         *uint                  `json:"limit,omitempty"`
	CommunityId   *uint                  `json:"community_id,omitempty"`
	CommunityName *string                `json:"community_name,omitempty"`
	PostId        *uint                  `json:"post_id,omitempty"`
	ParentId      *uint                  `json:"parent_id,omitempty"`
	SavedOnly     *bool                  `json:"saved_only,omitempty"`
}
