package lemmy

import "LemmyBeProxy/dto/model/lemmy"

type GetPersonMentionsResponse struct {
	Mentions []lemmy.PersonMentionView `json:"mentions" validate:"required"`
}
