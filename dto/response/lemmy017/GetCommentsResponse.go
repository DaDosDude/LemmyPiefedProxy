package lemmy017

import model017 "LemmyBeProxy/dto/model/lemmy017"

type GetCommentsResponse struct {
	Comments []model017.CommentView `json:"comments" validate:"required"`
}
