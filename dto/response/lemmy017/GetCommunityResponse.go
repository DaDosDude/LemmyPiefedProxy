package lemmy017

import (
	"LemmyBeProxy/dto/model/lemmy"
	model017 "LemmyBeProxy/dto/model/lemmy017"
)

// GetCommunityResponse matches Lemmy 0.17.2's real GetCommunityResponse.
// Site is always nil here (Option<Site> tolerates absence in Rust's
// deserialization) — we don't have a remote site's full data readily
// available in this flow. Online is 0, same honest-not-guessed reasoning
// as Post's online field. DiscussionLanguages is an empty (not nil)
// array — Piefed has no per-community language restriction concept to
// draw this from. DefaultPostLanguage is nil.
type GetCommunityResponse struct {
	CommunityView       model017.CommunityView         `json:"community_view" validate:"required"`
	Site                *lemmy.Site                    `json:"site,omitempty"`
	Moderators          []lemmy.CommunityModeratorView `json:"moderators" validate:"required"`
	Online              uint                           `json:"online" validate:"required"`
	DiscussionLanguages []uint                         `json:"discussion_languages" validate:"required"`
	DefaultPostLanguage *uint                          `json:"default_post_language,omitempty"`
}
