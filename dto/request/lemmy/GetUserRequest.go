package lemmy

import "LemmyPiefedApi/dto/model/lemmy"

// One of PersonId or Username must be set — same requirement Piefed enforces.
// CommunityId (limit posts/comments to one community) exists on Lemmy's
// GetPersonDetails but isn't wired through here yet.
type GetUserRequest struct {
	PersonId  *uint            `json:"person_id,omitempty"`
	Username  *string          `json:"username,omitempty"`
	Sort      *lemmy.SortType  `json:"sort,omitempty"`
	Page      *uint            `json:"page,omitempty"`
	Limit     *uint            `json:"limit,omitempty"`
	SavedOnly *bool            `json:"saved_only,omitempty"`
}
