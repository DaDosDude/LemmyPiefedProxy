package piefed

// Score: -1 to downvote, 1 to upvote, 0 to revert a previous vote.
// Not validate:"required" — 0 is meaningful here, not a missing field.
type LikeCommentRequest struct {
	CommentId uint `json:"comment_id" validate:"required"`
	Score     int  `json:"score"`
}
