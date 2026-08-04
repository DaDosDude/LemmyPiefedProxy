package lemmy

import "LemmyBeProxy/dto/model/lemmy"

type GetPostsRequest struct {
	CommunityId   *uint              `json:"community_id,omitempty"`
	CommunityName *string            `json:"community_name,omitempty"`
	DislikedOnly  *bool              `json:"disliked_only,omitempty"`
	LikedOnly     *bool              `json:"liked_only,omitempty"`
	Limit         *uint              `json:"limit,omitempty"`
	Page          *uint              `json:"page,omitempty"`
	PageCursor    *string            `json:"page_cursor,omitempty"`
	SavedOnly     *bool              `json:"saved_only,omitempty"`
	ShowHidden    *bool              `json:"show_hidden,omitempty"`
	ShowNsfw      *bool              `json:"show_nsfw,omitempty"`
	ShowRead      *bool              `json:"show_read,omitempty"`
	Sort          *lemmy.SortType    `json:"sort,omitempty"`
	Type          *lemmy.ListingType `json:"type_,omitempty"`
}
