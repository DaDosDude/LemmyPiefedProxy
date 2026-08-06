package lemmy017

import (
	"LemmyBeProxy/dto/model/lemmy"
	model017 "LemmyBeProxy/dto/model/lemmy017"
)

// GetUserResponse matches Lemmy 0.17.2's real GetPersonDetailsResponse.
// Comments and Posts use the 0.17-shaped nested views (same reasoning as
// GetCommentsResponse/GetPostsResponse) since they embed a creator —
// reusing canonical CommentView/PostView here would reintroduce the same
// nested-type gap already fixed for those endpoints directly.
type GetUserResponse struct {
	Comments   []model017.CommentView         `json:"comments" validate:"required"`
	Moderates  []lemmy.CommunityModeratorView `json:"moderates" validate:"required"`
	PersonView model017.PersonView            `json:"person_view" validate:"required"`
	Posts      []model017.PostView            `json:"posts" validate:"required"`
}

