package lemmy

import "LemmyBeProxy/dto/model/lemmy"

type GetCommunitiesRequest struct {
	Type     *lemmy.ListingType `json:"type_,omitempty"`
	Sort     *lemmy.SortType    `json:"sort,omitempty"`
	ShowNsfw *bool              `json:"show_nsfw,omitempty"`
	Page     *uint              `json:"page,omitempty"`
	Limit    *uint              `json:"limit,omitempty"`
}
