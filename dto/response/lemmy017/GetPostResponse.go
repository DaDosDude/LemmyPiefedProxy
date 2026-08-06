package lemmy017

import (
	model017 "LemmyBeProxy/dto/model/lemmy017"
)

// GetPostResponse matches Lemmy 0.17.2's real GetPostResponse: post_view,
// community_view, moderators, and online (a live viewer count for this
// post — later removed or changed in newer Lemmy versions). No
// cross_posts field — that's a later addition our canonical model
// includes but 0.17.x doesn't know about.
type GetPostResponse struct {
	PostView      model017.PostView                 `json:"post_view" required:"true"`
	CommunityView model017.CommunityView            `json:"community_view" required:"true"`
	Moderators    []model017.CommunityModeratorView `json:"moderators" required:"true"`
	Online        uint                              `json:"online" required:"true"`
}
