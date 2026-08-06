package lemmy017

import "LemmyBeProxy/dto/model/lemmy"

// CommunityView matches Lemmy 0.17.2's real CommunityView exactly:
// community, subscribed, blocked, counts. No embedded creator/person at
// all, unlike Post/CommentView — communities aren't authored by a single
// person the same way. Community is reused directly from canonical
// (confirmed field-by-field clean, same as it is for Post/CommentView).
type CommunityView struct {
	Community  lemmy.Community      `json:"community" validate:"required"`
	Subscribed lemmy.SubscribedType `json:"subscribed" validate:"required"`
	Blocked    bool                 `json:"blocked" validate:"required"`
	Counts     CommunityAggregates  `json:"counts" validate:"required"`
}
