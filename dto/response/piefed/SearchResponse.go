package piefed

import "LemmyBeProxy/dto/model/piefed"

type SearchResponse struct {
	Type        piefed.SearchType      `json:"type_" required:"true"`
	Communities []piefed.CommunityView `json:"communities" required:"true"`
	Posts       []piefed.PostView      `json:"posts" required:"true"`
	Users       []piefed.PersonView    `json:"users" required:"true"`
	Comments    []piefed.CommentView   `json:"comments" required:"true"`
}
