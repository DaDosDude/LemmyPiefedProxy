package lemmy

// Score: -1 to downvote, 1 to upvote, 0 to remove a previous vote.
// Not validate:"required" — 0 is meaningful here, not a missing field.
type CreateCommentLikeRequest struct {
	CommentId uint `json:"comment_id" validate:"required"`
	Score     int  `json:"score"`
}
