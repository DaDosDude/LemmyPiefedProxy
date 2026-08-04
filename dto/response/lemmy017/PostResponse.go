package lemmy017

import "LemmyBeProxy/dto/model/lemmy"

// PostResponse is what Lemmy 0.17.2 actually returns for post creation,
// editing, liking, and marking as read — confirmed against each of
// those four handlers in Lemmy's own source. Just post_view, none of the
// extra context GetPostResponse carries.
type PostResponse struct {
	PostView lemmy.PostView `json:"post_view" required:"true"`
}
