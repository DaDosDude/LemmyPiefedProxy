package lemmy017

// FormId is accepted here (so the request doesn't fail parsing) but not
// currently echoed back in the response — see CommentResponse.go for
// why. A known, documented gap, not a silent drop.
type CreateCommentRequest struct {
	Content    string  `json:"content" validate:"required"`
	PostId     uint    `json:"post_id" validate:"required"`
	ParentId   *uint   `json:"parent_id,omitempty"`
	LanguageId *uint   `json:"language_id,omitempty"`
	FormId     *string `json:"form_id,omitempty"`
}
