package lemmy

// Upvotes and Downvotes are int, not uint — matches go-lemmy's own
// PostAggregates (mlmym's client library), which uses int64 here. A real
// Piefed response was observed with a negative downvotes value.
type PostAggregates struct {
	Comments          uint   `json:"comments" validate:"required"`
	Downvotes         int    `json:"downvotes" validate:"required"`
	NewestCommentTime string `json:"newest_comment_time" validate:"required"`
	PostId            uint   `json:"post_id" validate:"required"`
	Published         string `json:"published" validate:"required"`
	Score             int    `json:"score" validate:"required"`
	Upvotes           int    `json:"upvotes" validate:"required"`
}
