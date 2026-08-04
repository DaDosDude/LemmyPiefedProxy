package lemmy017

// CommentAggregates matches Lemmy 0.17.2's real CommentAggregates.
// Confirmed missing from our canonical CommentAggregates: Id (the
// aggregates row's own primary key, distinct from CommentId — defaulted
// to CommentId's value, same harmless stand-in used for PostAggregates).
type CommentAggregates struct {
	Id         uint   `json:"id" validate:"required"`
	CommentId  uint   `json:"comment_id" validate:"required"`
	Score      int    `json:"score" validate:"required"`
	Upvotes    int    `json:"upvotes" validate:"required"`
	Downvotes  int    `json:"downvotes" validate:"required"`
	Published  string `json:"published" validate:"required"`
	ChildCount uint   `json:"child_count" validate:"required"`
}
