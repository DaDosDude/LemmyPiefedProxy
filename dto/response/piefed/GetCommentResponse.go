package piefed

import (
	"LemmyBeProxy/dto/model/piefed"
)

type GetCommentResponse struct {
	CommentView piefed.CommentView `json:"comment_view" validate:"required"`
}
