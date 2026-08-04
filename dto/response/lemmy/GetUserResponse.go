package lemmy

import "LemmyBeProxy/dto/model/lemmy"

// Lemmy's GetPersonDetailsResponse also has an optional Site field
// (ConvertSite requires a separate ActivityPub actor fetch, same as
// SiteController.Site() does) — not wired here, so it's simply absent
// rather than guessed at.
type GetUserResponse struct {
	Comments   []lemmy.CommentView            `json:"comments" required:"true"`
	Moderates  []lemmy.CommunityModeratorView `json:"moderates" required:"true"`
	PersonView lemmy.PersonView               `json:"person_view" required:"true"`
	Posts      []lemmy.PostView               `json:"posts" required:"true"`
}
