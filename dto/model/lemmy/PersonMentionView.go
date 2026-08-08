package lemmy

type PersonMentionView struct {
	PersonMention              PersonMention     `json:"person_mention" validate:"required"`
	Comment                    Comment           `json:"comment" validate:"required"`
	Creator                    Person            `json:"creator" validate:"required"`
	Post                       Post              `json:"post" validate:"required"`
	Community                  Community         `json:"community" validate:"required"`
	Recipient                  Person            `json:"recipient" validate:"required"`
	Counts                     CommentAggregates `json:"counts" validate:"required"`
	CreatorBannedFromCommunity bool              `json:"creator_banned_from_community" validate:"required"`
	Subscribed                 SubscribedType    `json:"subscribed" validate:"required"`
	Saved                      bool              `json:"saved" validate:"required"`
	CreatorBlocked             bool              `json:"creator_blocked" validate:"required"`
	MyVote                     *int              `json:"my_vote,omitempty"`
}
