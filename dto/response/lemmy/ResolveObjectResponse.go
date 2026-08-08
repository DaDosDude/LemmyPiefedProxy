package lemmy

import "LemmyBeProxy/dto/model/lemmy"

// ResolveObjectResponse matches real Lemmy's shape exactly — a union
// type where only one of these four is ever populated, depending on
// what kind of object the query string resolved to.
type ResolveObjectResponse struct {
	Comment   *lemmy.CommentView   `json:"comment,omitempty"`
	Post      *lemmy.PostView      `json:"post,omitempty"`
	Community *lemmy.CommunityView `json:"community,omitempty"`
	Person    *lemmy.PersonView    `json:"person,omitempty"`
}
