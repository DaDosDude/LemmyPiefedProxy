package lemmy017

import "LemmyBeProxy/dto/model/lemmy"

type CommunityModeratorView struct {
	Community lemmy.Community `json:"community" validate:"required"`
	Moderator Person          `json:"moderator" validate:"required"`
}
