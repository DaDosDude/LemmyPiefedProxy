package lemmy

import "LemmyBeProxy/dto/model/lemmy"

type SearchResponse struct {
	Type        lemmy.SearchType      `json:"type_" required:"true"`
	Communities []lemmy.CommunityView `json:"communities" required:"true"`
	Posts       []lemmy.PostView      `json:"posts" required:"true"`
	Users       []lemmy.PersonView    `json:"users" required:"true"`
	Comments    []lemmy.CommentView   `json:"comments" required:"true"`
}
