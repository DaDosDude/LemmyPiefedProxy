package lemmy017

import model017 "LemmyBeProxy/dto/model/lemmy017"

type GetPersonMentionsResponse struct {
	Mentions []model017.PersonMentionView `json:"mentions" validate:"required"`
}
