package lemmy017

import "LemmyBeProxy/dto/model/lemmy"

// GetPostsResponse matches Lemmy 0.17.2's real GetPostsResponse — just
// posts, no next_page. Cursor-based pagination didn't exist yet.
type GetPostsResponse struct {
	Posts []lemmy.PostView `json:"posts" required:"true"`
}
