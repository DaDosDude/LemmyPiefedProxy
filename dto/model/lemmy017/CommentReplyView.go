package lemmy017

import "LemmyBeProxy/dto/model/lemmy"

type CommentReplyView struct {
	CommentReply               lemmy.CommentReply  `json:"comment_reply" validate:"required"`
	Comment                    lemmy.Comment       `json:"comment" validate:"required"`
	Creator                    Person              `json:"creator" validate:"required"`
	Post                       lemmy.Post          `json:"post" validate:"required"`
	Community                  lemmy.Community     `json:"community" validate:"required"`
	Recipient                  Person              `json:"recipient" validate:"required"`
	Counts                     CommentAggregates   `json:"counts" validate:"required"`
	CreatorBannedFromCommunity bool                `json:"creator_banned_from_community" validate:"required"`
	Subscribed                 lemmy.SubscribedType `json:"subscribed" validate:"required"`
	Saved                      bool                `json:"saved" validate:"required"`
	CreatorBlocked             bool                `json:"creator_blocked" validate:"required"`
	MyVote                     *int                `json:"my_vote" validate:"required"`
}
