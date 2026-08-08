package lemmy

import "LemmyBeProxy/dto/model/lemmy"

type GetRepliesResponse struct {
	Replies []lemmy.CommentReplyView `json:"replies" validate:"required"`
}
