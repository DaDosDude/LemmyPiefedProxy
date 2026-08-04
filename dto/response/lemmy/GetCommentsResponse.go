package lemmy

import "LemmyBeProxy/dto/model/lemmy"

type GetCommentsResponse struct {
	Comments []lemmy.CommentView `json:"comments" required:"true"`
}
