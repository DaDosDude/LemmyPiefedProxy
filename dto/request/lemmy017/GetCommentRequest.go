package lemmy017

type GetCommentRequest struct {
	Id uint `json:"id" validate:"required"`
}
