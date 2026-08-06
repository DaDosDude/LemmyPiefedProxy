package lemmy017

import "LemmyBeProxy/dto/model/lemmy"

type CommunityFollowerView struct {
	Community lemmy.Community `json:"community" validate:"required"`
	Follower  Person          `json:"follower" validate:"required"`
}
