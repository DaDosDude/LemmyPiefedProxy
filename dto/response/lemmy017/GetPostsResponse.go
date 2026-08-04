package lemmy017

import model017 "LemmyBeProxy/dto/model/lemmy017"

// GetPostsResponse matches Lemmy 0.17.2's real GetPostsResponse — just
// posts, no next_page. Cursor-based pagination didn't exist yet.
type GetPostsResponse struct {
	Posts []model017.PostView `json:"posts" required:"true"`
}
