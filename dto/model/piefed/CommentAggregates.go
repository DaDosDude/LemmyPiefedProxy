package piefed

// Upvotes/Downvotes are int, not uint — see PostAggregates.go for why:
// a real Piefed response can carry a negative value here.
type CommentAggregates struct {
	CommentId  uint   `json:"comment_id" validate:"required"`
	Score      int    `json:"score" validate:"required"`
	Upvotes    int    `json:"upvotes" validate:"required"`
	Downvotes  int    `json:"downvotes" validate:"required"`
	Published  string `json:"published" validate:"required"`
	ChildCount uint   `json:"child_count" validate:"required"`
}
