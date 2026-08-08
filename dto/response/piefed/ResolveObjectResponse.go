package piefed

import "LemmyBeProxy/dto/model/piefed"

type ResolveObjectResponse struct {
	Comment   *piefed.CommentView   `json:"comment,omitempty"`
	Post      *piefed.PostView      `json:"post,omitempty"`
	Community *piefed.CommunityView `json:"community,omitempty"`
	Person    *piefed.PersonView    `json:"person,omitempty"`
}
