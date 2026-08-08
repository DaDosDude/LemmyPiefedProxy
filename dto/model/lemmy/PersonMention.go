package lemmy

type PersonMention struct {
	Id          uint   `json:"id" validate:"required"`
	RecipientId uint   `json:"recipient_id" validate:"required"`
	CommentId   uint   `json:"comment_id" validate:"required"`
	Read        bool   `json:"read"`
	Published   string `json:"published" validate:"required"`
}
