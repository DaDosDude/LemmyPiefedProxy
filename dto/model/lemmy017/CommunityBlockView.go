package lemmy017

import "LemmyBeProxy/dto/model/lemmy"

type CommunityBlockView struct {
	Person    Person          `json:"person" validate:"required"`
	Community lemmy.Community `json:"community" validate:"required"`
}
