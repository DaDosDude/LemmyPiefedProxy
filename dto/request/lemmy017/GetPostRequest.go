package lemmy017

type GetPostRequest struct {
	Id        *uint `json:"id,omitempty"`
	CommentId *uint `json:"comment_id,omitempty"`
}
