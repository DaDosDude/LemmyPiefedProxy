package lemmy017

import "LemmyBeProxy/dto/model/lemmy"

// GetPostsRequest matches Lemmy 0.17.2's real GetPosts struct exactly
// (confirmed against crates/api_common/src/post.rs in Lemmy's own source
// at tag 0.17.2). Notably: no page_cursor (0.17.x has no cursor
// pagination concept at all), no liked_only/disliked_only/show_hidden/
// show_read (all later additions). auth is not modeled here — it's
// extracted centrally before routing, the same mechanism already built
// for general Lemmy 0.18.x compatibility.
type GetPostsRequest struct {
	Type          *lemmy.ListingType `json:"type_,omitempty"`
	Sort          *lemmy.SortType    `json:"sort,omitempty"`
	Page          *uint              `json:"page,omitempty"`
	Limit         *uint              `json:"limit,omitempty"`
	CommunityId   *uint              `json:"community_id,omitempty"`
	CommunityName *string            `json:"community_name,omitempty"`
	SavedOnly     *bool              `json:"saved_only,omitempty"`
}
