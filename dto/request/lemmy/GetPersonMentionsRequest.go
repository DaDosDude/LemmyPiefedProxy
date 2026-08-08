package lemmy

import "LemmyBeProxy/dto/model/lemmy"

type GetPersonMentionsRequest struct {
	Sort       *lemmy.CommentSortType `json:"sort,omitempty"`
	Page       *uint                  `json:"page,omitempty"`
	Limit      *uint                  `json:"limit,omitempty"`
	UnreadOnly *bool                  `json:"unread_only,omitempty"`
}
