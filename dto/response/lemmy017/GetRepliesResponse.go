package lemmy017

import model017 "LemmyBeProxy/dto/model/lemmy017"

type GetRepliesResponse struct {
	Replies []model017.CommentReplyView `json:"replies" validate:"required"`
}
