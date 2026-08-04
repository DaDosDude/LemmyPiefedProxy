package piefed

import "LemmyBeProxy/dto/model/piefed"

type GetCommunitiesRequest struct {
	Type     *piefed.ListingType `json:"type_,omitempty"`
	Sort     *piefed.SortType    `json:"sort,omitempty"`
	ShowNsfw *bool               `json:"show_nsfw,omitempty"`
	Page     *uint               `json:"page,omitempty"`
	Limit    *uint               `json:"limit,omitempty"`
}
