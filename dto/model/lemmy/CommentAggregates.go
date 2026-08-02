package lemmy

// Upvotes/Downvotes are int, not uint — matches go-lemmy's own
// CommentAggregates (int64). See PostAggregates.go for the full reasoning.
type CommentAggregates struct {
	ChildCount uint   `json:"child_count" validate:"required"`
	CommentId  uint   `json:"comment_id" validate:"required"`
	Downvotes  int    `json:"downvotes" validate:"required"`
	Published  string `json:"published" validate:"required"`
	Score      int    `json:"score" validate:"required"`
	Upvotes    int    `json:"upvotes" validate:"required"`
}
