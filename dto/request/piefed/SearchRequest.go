package piefed

import "LemmyPiefedApi/dto/model/piefed"

type SearchRequest struct {
	Q             string              `json:"q"`
	Type          piefed.SearchType   `json:"type_"`
	Limit         *uint               `json:"limit,omitempty"`
	ListingType   *piefed.ListingType `json:"listing_type,omitempty"`
	Page          *uint               `json:"page,omitempty"`
	Sort          *piefed.SortType    `json:"sort,omitempty"`
	CommunityName *string             `json:"community_name,omitempty"`
	CommunityId   *uint               `json:"community_id,omitempty"`
}
