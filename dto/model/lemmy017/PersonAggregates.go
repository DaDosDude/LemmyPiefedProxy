package lemmy017

// PersonAggregates matches Lemmy 0.17.2's real PersonAggregates. Same
// extra-id pattern as other Aggregates types, plus PostScore and
// CommentScore — fields our canonical model never tracked at all, not
// just omitted for 0.17.x. Both default to 0, an honest "we don't have
// this data" rather than a guess.
type PersonAggregates struct {
	Id           uint `json:"id" validate:"required"`
	PersonId     uint `json:"person_id" validate:"required"`
	PostCount    uint `json:"post_count" validate:"required"`
	PostScore    int  `json:"post_score" validate:"required"`
	CommentCount uint `json:"comment_count" validate:"required"`
	CommentScore int  `json:"comment_score" validate:"required"`
}
