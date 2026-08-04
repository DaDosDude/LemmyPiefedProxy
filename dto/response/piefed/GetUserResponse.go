package piefed

import "LemmyBeProxy/dto/model/piefed"

type GetUserResponse struct {
	Comments   []piefed.CommentView            `json:"comments" required:"true"`
	Moderates  []piefed.CommunityModeratorView `json:"moderates" required:"true"`
	PersonView piefed.PersonView               `json:"person_view" required:"true"`
	Posts      []piefed.PostView               `json:"posts" required:"true"`
}
