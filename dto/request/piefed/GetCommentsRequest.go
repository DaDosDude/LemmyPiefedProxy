package piefed

import (
	"LemmyBeProxy/dto/model/piefed"
)

type GetCommentsRequest struct {
	Type        *piefed.ListingType     `json:"type_,omitempty"`
	PersonId    *uint                   `json:"person_id,omitempty"`
	MaxDepth    *uint                   `json:"max_depth,omitempty"`
	Page        *uint                   `json:"page,omitempty"`
	ParentId    *uint                   `json:"parent_id,omitempty"`
	CommunityId *uint                   `json:"community_id,omitempty"`
	PostId      *uint                   `json:"post_id,omitempty"`
	Limit       *uint                   `json:"limit,omitempty"`
	Sort        *piefed.CommentSortType `json:"sort,omitempty"`
}
