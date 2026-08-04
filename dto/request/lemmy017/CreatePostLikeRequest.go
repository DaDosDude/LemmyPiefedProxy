package lemmy017

// Score is deliberately NOT validate:"required" — 0 is a legitimate
// value (removing a vote), and go-playground/validator's "required" tag
// rejects the zero value on plain int fields.
type CreatePostLikeRequest struct {
	PostId uint `json:"post_id" validate:"required"`
	Score  int  `json:"score"`
}
