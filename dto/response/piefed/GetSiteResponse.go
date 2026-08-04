package piefed

import "LemmyBeProxy/dto/model/piefed"

type GetSiteResponse struct {
	MyUser  *piefed.MyUserInfo  `json:"my_user,omitempty"`
	Site    piefed.Site         `json:"site" validate:"required"`
	Version string              `json:"version" validate:"required"`
	Admins  []piefed.PersonView `json:"admins" validate:"required"`
}
