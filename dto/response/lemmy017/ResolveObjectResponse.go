package lemmy017

import model017 "LemmyBeProxy/dto/model/lemmy017"

// ResolveObjectResponse matches Lemmy 0.17.2's real shape — a union
// type where only one field is ever populated. No omitempty: real
// 0.17.x's deserializer requires the key present (as null) even when
// the value is absent, same lesson as everywhere else in this package.
type ResolveObjectResponse struct {
	Comment   *model017.CommentView   `json:"comment"`
	Post      *model017.PostView      `json:"post"`
	Community *model017.CommunityView `json:"community"`
	Person    *model017.PersonView    `json:"person"`
}
