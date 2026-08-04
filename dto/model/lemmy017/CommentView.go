package lemmy017

import "LemmyBeProxy/dto/model/lemmy"

// CommentView matches Lemmy 0.17.2's real CommentView field set exactly
// — confirmed against Lemmy's own source. Comment, Post, and Community
// are reused directly from the canonical model since all three were
// checked field-by-field against 0.17.2 and found to already contain
// every field 0.17.x requires. Creator and Counts use the 0.17-specific
// types in this package instead, since those two genuinely differ (same
// split as PostView).
type CommentView struct {
	Comment                    lemmy.Comment        `json:"comment" validate:"required"`
	Creator                    Person               `json:"creator" validate:"required"`
	Post                       lemmy.Post           `json:"post" validate:"required"`
	Community                  lemmy.Community      `json:"community" validate:"required"`
	Counts                     CommentAggregates    `json:"counts" validate:"required"`
	CreatorBannedFromCommunity bool                 `json:"creator_banned_from_community" validate:"required"`
	Subscribed                 lemmy.SubscribedType `json:"subscribed" validate:"required"`
	Saved                      bool                 `json:"saved" validate:"required"`
	CreatorBlocked             bool                 `json:"creator_blocked" validate:"required"`
	MyVote                     *int                 `json:"my_vote,omitempty"`
}
