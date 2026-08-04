package lemmy017

// Score is deliberately NOT validate:"required" — 0 is a legitimate
// value (removing a vote), same reasoning as CreatePostLikeRequest.
type CreateCommentLikeRequest struct {
	CommentId uint `json:"comment_id" validate:"required"`
	Score     int  `json:"score"`
}
