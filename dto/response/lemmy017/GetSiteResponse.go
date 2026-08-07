package lemmy017

import (
	"LemmyBeProxy/dto/model/lemmy"
	model017 "LemmyBeProxy/dto/model/lemmy017"
)

// GetSiteResponse matches Lemmy 0.17.2's real GetSiteResponse. Online is
// 0, same honest-not-guessed reasoning as elsewhere. FederatedInstances
// is always nil here (Option<FederatedInstances> tolerates absence) —
// our canonical model has no equivalent type at all, and mapping
// Piefed's federation data into this shape hasn't been done, a separate
// piece of work. AllLanguages and Taglines reuse canonical directly —
// both confirmed field-by-field clean.
type GetSiteResponse struct {
	SiteView            model017.SiteView     `json:"site_view" validate:"required"`
	Admins              []model017.PersonView `json:"admins" validate:"required"`
	Online              uint                  `json:"online" validate:"required"`
	Version             string                `json:"version" validate:"required"`
	MyUser              *model017.MyUserInfo  `json:"my_user"`
	AllLanguages        []lemmy.Language      `json:"all_languages" validate:"required"`
	DiscussionLanguages []uint                `json:"discussion_languages" validate:"required"`
	Taglines            []lemmy.Tagline       `json:"taglines"`
	FederatedInstances  *struct{}             `json:"federated_instances"`
}

