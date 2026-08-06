package lemmy017

import model017 "LemmyBeProxy/dto/model/lemmy017"

// SearchResponse matches Lemmy 0.17.2's real SearchResponse. Every
// result array uses the 0.17-shaped nested view types, same reasoning
// as GetUserResponse's Comments/Posts — each embeds a creator, so
// reusing canonical types here would reintroduce the same nested-type
// gap already fixed elsewhere.
type SearchResponse struct {
	Type        string                   `json:"type_" validate:"required"`
	Comments    []model017.CommentView   `json:"comments" validate:"required"`
	Posts       []model017.PostView      `json:"posts" validate:"required"`
	Communities []model017.CommunityView `json:"communities" validate:"required"`
	Users       []model017.PersonView    `json:"users" validate:"required"`
}
