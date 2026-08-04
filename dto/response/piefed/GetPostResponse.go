package piefed

import "LemmyBeProxy/dto/model/piefed"

type GetPostResponse struct {
	PostView      piefed.PostView                 `json:"post_view" required:"true"`
	CommunityView piefed.CommunityView            `json:"community_view" required:"true"`
	Moderators    []piefed.CommunityModeratorView `json:"moderators" required:"true"`
	CrossPosts    []piefed.PostView               `json:"cross_posts" required:"true"`
}
