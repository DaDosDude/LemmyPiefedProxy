package piefed

// Score: -1 to downvote, 1 to upvote, 0 to revert a previous vote.
// Deliberately NOT tagged validate:"required" — 0 is a legitimate,
// meaningful value here (unvote), and go-playground/validator's
// "required" tag rejects the zero value on plain int fields.
type LikePostRequest struct {
	PostId uint `json:"post_id" validate:"required"`
	Score  int  `json:"score"`
}
