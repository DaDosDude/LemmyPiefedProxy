package lemmy017

import "LemmyBeProxy/dto/model/lemmy"

// GetCommunitiesRequest matches Lemmy 0.17.2's real ListCommunities.
// No show_nsfw field at all — confirmed against Lemmy's own source,
// unlike our canonical model which has it as a later addition.
type GetCommunitiesRequest struct {
	Type  *lemmy.ListingType `json:"type_,omitempty"`
	Sort  *lemmy.SortType    `json:"sort,omitempty"`
	Page  *uint              `json:"page,omitempty"`
	Limit *uint              `json:"limit,omitempty"`
}
