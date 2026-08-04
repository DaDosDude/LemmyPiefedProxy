package lemmy

import "LemmyBeProxy/dto/model/lemmy"

type GetCommentResponse struct {
	CommentView  lemmy.CommentView `json:"comment_view" validate:"required"`
	RecipientIds []uint            `json:"recipient_ids" validate:"required"`
}
