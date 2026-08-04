package lemmy

import "LemmyBeProxy/dto/model/lemmy"

type GetPostResponse struct {
	CommunityView lemmy.CommunityView            `json:"community_view" required:"true"`
	CrossPosts    []lemmy.PostView               `json:"cross_posts" required:"true"`
	Moderators    []lemmy.CommunityModeratorView `json:"moderators" required:"true"`
	PostView      lemmy.PostView                 `json:"post_view" required:"true"`
}
