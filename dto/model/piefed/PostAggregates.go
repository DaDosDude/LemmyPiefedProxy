package piefed

// Upvotes and Downvotes are int, not uint — a real Piefed response was
// observed with downvotes: -1 (an apparent data anomaly on Piefed's own
// side), and json.Unmarshal hard-fails the entire response the moment it
// hits a negative number on a uint field. int tolerates it; Score was
// already int for the same reason.
type PostAggregates struct {
	PostId            uint   `json:"post_id" validate:"required"`
	Comments          uint   `json:"comments" validate:"required"`
	Score             int    `json:"score" validate:"required"`
	Upvotes           int    `json:"upvotes" validate:"required"`
	Downvotes         int    `json:"downvotes" validate:"required"`
	Published         string `json:"published" validate:"required"`
	NewestCommentTime string `json:"newest_comment_time" validate:"required"`
}
